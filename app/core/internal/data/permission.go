package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	permissionv1 "github.com/realHoangHai/kratos/api/gen/go/permission/v1"
	"github.com/realHoangHai/kratos/app/core/internal/data/ent"
	"github.com/realHoangHai/kratos/app/core/internal/data/ent/permission"
	util "github.com/realHoangHai/kratos/pkg/utils/time"
	"github.com/realHoangHai/kratos/pkg/utils/trans"
)

type PermissionRepo struct {
	data *Data
	log  *log.Helper
}

func NewPermissionRepo(data *Data, logger log.Logger) *PermissionRepo {
	l := log.NewHelper(log.With(logger, "module", "permission/data/core-service"))
	return &PermissionRepo{
		data: data,
		log:  l,
	}
}

func (r *PermissionRepo) convertEntToProto(data *ent.Permission) *permissionv1.Permission {
	if data == nil {
		return nil
	}

	return &permissionv1.Permission{
		Id:          data.ID,
		Name:        trans.String(data.Name),
		GuardName:   trans.String(data.GuardName),
		Description: data.Description,
		CreateTime:  util.UnixMilliToStringPtr(data.CreateTime),
		UpdateTime:  util.UnixMilliToStringPtr(data.UpdateTime),
		CreatorId:   nil,
	}
}

func (r *PermissionRepo) GetByID(ctx context.Context, req *permissionv1.GetPermissionRequest) (*permissionv1.Permission, error) {
	ret, err := r.data.db.Client().Permission.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("Query one data failed: %s", err.Error())
		return nil, err
	}

	p := r.convertEntToProto(ret)

	return p, err
}

func (r *PermissionRepo) GetByGuardName(ctx context.Context, guardName string) (*permissionv1.Permission, error) {
	ret, err := r.data.db.Client().Permission.Query().
		Where(permission.GuardNameEQ(guardName)).
		Only(ctx)
	if err != nil {
		r.log.Errorf("Query permission data failed: %s", err.Error())
		return nil, err
	}

	p := r.convertEntToProto(ret)

	return p, err
}
