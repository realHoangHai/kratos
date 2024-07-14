package service

import (
	"context"
	adminv1 "github.com/realHoangHai/kratos/api/gen/go/admin/v1"
	commonv1 "github.com/realHoangHai/kratos/api/gen/go/common/v1"
	userv1 "github.com/realHoangHai/kratos/api/gen/go/user/v1"
	"github.com/realHoangHai/kratos/handler/middleware/auth"
	"github.com/realHoangHai/kratos/pkg/utils/trans"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	adminv1.UserServiceHTTPServer

	uc  userv1.UserServiceClient
	log *log.Helper
}

func NewUserService(logger log.Logger, uc userv1.UserServiceClient) *UserService {
	l := log.NewHelper(log.With(logger, "module", "admin/service/user"))
	return &UserService{
		log: l,
		uc:  uc,
	}
}

func (s *UserService) ListUser(ctx context.Context, req *commonv1.PagingRequest) (*userv1.ListUserResponse, error) {
	return s.uc.ListUser(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.User, error) {
	return s.uc.GetUser(ctx, req)
}

func (s *UserService) GetUserByUserName(ctx context.Context, req *userv1.GetUserByUserNameRequest) (*userv1.User, error) {
	return s.uc.GetUserByUserName(ctx, req)
}

func (s *UserService) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.User, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("User authentication failed [%s]", err.Error())
		return nil, adminv1.ErrorAccessForbidden("User authentication failed")
	}

	if req.User == nil {
		return nil, adminv1.ErrorBadRequest("Wrong parameter")
	}

	req.OperatorId = authInfo.UserId
	req.User.CreatorId = trans.Uint32(authInfo.UserId)
	if req.User.Authority == nil {
		req.User.Authority = userv1.UserAuthority_CUSTOMER_USER.Enum()
	}

	return s.uc.CreateUser(ctx, req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.User, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("User authentication failed [%s]", err.Error())
		return nil, adminv1.ErrorAccessForbidden("User authentication failed")
	}

	if req.User == nil {
		return nil, adminv1.ErrorBadRequest("Wrong parameter")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.UpdateUser(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *userv1.DeleteUserRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("User authentication failed [%s]", err.Error())
		return nil, adminv1.ErrorAccessForbidden("User authentication failed")
	}

	req.OperatorId = authInfo.UserId

	return s.uc.DeleteUser(ctx, req)
}

func (s *UserService) UserExists(ctx context.Context, req *userv1.UserExistsRequest) (*userv1.UserExistsResponse, error) {
	return s.uc.UserExists(ctx, req)
}
