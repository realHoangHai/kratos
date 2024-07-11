package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	commonv1 "github.com/realHoangHai/kratos/api/gen/go/common/v1"
	userv1 "github.com/realHoangHai/kratos/api/gen/go/user/v1"
	"github.com/realHoangHai/kratos/internal/data"
	"github.com/realHoangHai/kratos/pkg/utils/trans"
)

type UserService struct {
	userv1.UnimplementedUserServiceServer

	uc  *data.UserRepo
	log *log.Helper
}

func NewUserService(logger log.Logger, uc *data.UserRepo) *UserService {
	l := log.NewHelper(log.With(logger, "module", "service/user"))
	return &UserService{
		uc:  uc,
		log: l,
	}
}

func (s *UserService) ListUser(ctx context.Context, req *commonv1.PagingRequest) (*userv1.ListUserResponse, error) {
	resp, err := s.uc.List(ctx, req)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(resp.Items); i++ {
		resp.Items[i].Password = trans.String("")
	}

	return resp, nil
}

func (s *UserService) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.User, error) {
	resp, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	resp.Password = trans.String("")

	return resp, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.User, error) {
	resp, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	resp.Password = trans.String("")

	return resp, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.User, error) {
	resp, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	resp.Password = trans.String("")

	return resp, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *userv1.DeleteUserRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
