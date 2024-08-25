package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/go-kratos/kratos/v2/log"

	commonv1 "github.com/realHoangHai/kratos/api/gen/go/common/v1"
	userv1 "github.com/realHoangHai/kratos/api/gen/go/user/v1"
	"github.com/realHoangHai/kratos/app/core/internal/data/ent"
	"github.com/realHoangHai/kratos/app/core/internal/data/ent/user"
	"github.com/realHoangHai/kratos/pkg/utils/crypto"
	entgo "github.com/realHoangHai/kratos/pkg/utils/entgo/query"
	util "github.com/realHoangHai/kratos/pkg/utils/time"
	"github.com/realHoangHai/kratos/pkg/utils/trans"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/data/core-service"))
	return &UserRepo{
		data: data,
		log:  l,
	}
}

func (r *UserRepo) convertEntToProto(data *ent.User) *userv1.User {
	if data == nil {
		return nil
	}

	var authority *userv1.UserAuthority
	if data.Authority != nil {
		authority = (*userv1.UserAuthority)(trans.Int32(userv1.UserAuthority_value[string(*data.Authority)]))
	}

	return &userv1.User{
		Id:          data.ID,
		UserName:    data.Username,
		Email:       data.Email,
		Avatar:      data.Avatar,
		Description: data.Description,
		Password:    data.Password,
		Authority:   authority,
		CreateTime:  util.UnixMilliToStringPtr(data.CreateTime),
		UpdateTime:  util.UnixMilliToStringPtr(data.UpdateTime),
	}
}

func (r *UserRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().User.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("Query count failed: %s", err.Error())
	}

	return count, err
}

func (r *UserRepo) List(ctx context.Context, req *commonv1.PagingRequest) (*userv1.ListUserResponse, error) {
	builder := r.data.db.Client().User.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), user.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("An error occurred in parsing conditions [%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("Query list failed: %s", err.Error())
		return nil, err
	}

	items := make([]*userv1.User, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		if item != nil {
			item.Password = nil
		}
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &userv1.ListUserResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *UserRepo) Get(ctx context.Context, req *userv1.GetUserRequest) (*userv1.User, error) {
	ret, err := r.data.db.Client().User.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("Query one data failed: %s", err.Error())
		return nil, err
	}

	u := r.convertEntToProto(ret)
	if u != nil {
		u.Password = nil
	}

	return u, err
}

func (r *UserRepo) Create(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.User, error) {
	ph, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	builder := r.data.db.Client().User.Create().
		SetNillableUsername(req.User.UserName).
		SetNillableEmail(req.User.Email).
		SetPassword(ph).
		SetCreateTime(time.Now().UnixMilli())

	if req.User.Authority != nil {
		builder.SetAuthority((user.Authority)(req.User.Authority.String()))
	}

	ret, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("Insert one data failed: %s", err.Error())
		return nil, err
	}

	u := r.convertEntToProto(ret)
	if u != nil {
		u.Password = nil
	}

	return u, err
}

func (r *UserRepo) Update(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.User, error) {
	cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	builder := r.data.db.Client().User.UpdateOneID(req.Id).
		SetNillableEmail(req.User.Email).
		SetPassword(cryptoPassword).
		SetUpdateTime(time.Now().UnixMilli())

	if req.User.Authority != nil {
		builder.SetAuthority((user.Authority)(req.User.Authority.String()))
	}

	ret, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("Update one data failed: %s", err.Error())
		return nil, err
	}

	u := r.convertEntToProto(ret)
	if u != nil {
		u.Password = nil
	}

	return u, err
}

func (r *UserRepo) Delete(ctx context.Context, req *userv1.DeleteUserRequest) (bool, error) {
	err := r.data.db.Client().User.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("Delete one data failed: %s", err.Error())
	}

	return err == nil, err
}

func (r *UserRepo) GetUserByUserName(ctx context.Context, userName string) (*userv1.User, error) {
	ret, err := r.data.db.Client().User.Query().
		Where(user.UsernameEQ(userName)).
		Only(ctx)
	if err != nil {
		r.log.Errorf("Query user data failed: %s", err.Error())
		return nil, err
	}

	u := r.convertEntToProto(ret)
	if u != nil {
		u.Password = nil
	}

	return u, err
}

func (r *UserRepo) VerifyPassword(ctx context.Context, req *userv1.VerifyPasswordRequest) (bool, error) {
	res, err := r.data.db.Client().User.
		Query().
		Select(user.FieldID, user.FieldPassword).
		Where(user.UsernameEQ(req.GetUserName())).
		Only(ctx)

	if err != nil {
		r.log.Errorf("Query user data failed: %s", err.Error())
		return false, userv1.ErrorUserNotFound("User not found")
	}

	matched := crypto.CheckPasswordHash(req.GetPassword(), *res.Password)

	if !matched {
		return false, userv1.ErrorUserNotFound("Wrong password")
	}

	return true, nil
}
