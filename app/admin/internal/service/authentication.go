package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-kratos/kratos/v2/log"

	adminv1 "github.com/realHoangHai/kratos/api/gen/go/admin/v1"
	userv1 "github.com/realHoangHai/kratos/api/gen/go/user/v1"
	"github.com/realHoangHai/kratos/handler/cache"
	"github.com/realHoangHai/kratos/handler/middleware/auth"
)

type AuthenticationService struct {
	adminv1.AuthenticationServiceHTTPServer

	uc   userv1.UserServiceClient
	utuc *cache.UserToken

	log *log.Helper
}

func NewAuthenticationService(logger log.Logger, uc userv1.UserServiceClient, utuc *cache.UserToken) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "authentication/server/admin-service"))
	return &AuthenticationService{
		log:  l,
		uc:   uc,
		utuc: utuc,
	}
}

// Login .
func (s *AuthenticationService) Login(ctx context.Context, req *adminv1.LoginRequest) (*adminv1.LoginResponse, error) {
	if _, err := s.uc.VerifyPassword(ctx, &userv1.VerifyPasswordRequest{
		UserName: req.GetUserName(),
		Password: req.GetPassword(),
	}); err != nil {
		return &adminv1.LoginResponse{}, err
	}

	user, err := s.uc.GetUserByUserName(ctx, &userv1.GetUserByUserNameRequest{UserName: req.GetUserName()})
	if err != nil {
		return &adminv1.LoginResponse{}, err
	}

	accessToken, refreshToken, err := s.utuc.GenerateToken(ctx, user)
	if err != nil {
		return nil, err
	}

	return &adminv1.LoginResponse{
		TokenType:    adminv1.TokenType_bearer.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Logout 登出
func (s *AuthenticationService) Logout(ctx context.Context, req *adminv1.LogoutRequest) (*emptypb.Empty, error) {
	err := s.utuc.RemoveToken(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AuthenticationService) GetMe(ctx context.Context, req *adminv1.GetMeRequest) (*userv1.User, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("User authentication failed [%s]", err.Error())
		return nil, adminv1.ErrorAccessForbidden("User authentication failed")
	}

	req.Id = authInfo.UserId

	return s.uc.GetUser(ctx, &userv1.GetUserRequest{
		Id: req.GetId(),
	})
}
