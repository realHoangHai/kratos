package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-kratos/kratos/v2/log"

	clientv1 "github.com/realHoangHai/kratos/api/gen/go/client/v1"
	userv1 "github.com/realHoangHai/kratos/api/gen/go/user/v1"
	"github.com/realHoangHai/kratos/handler/cache"
	"github.com/realHoangHai/kratos/handler/middleware/auth"
)

type AuthenticationService struct {
	clientv1.AuthenticationServiceHTTPServer

	sc   userv1.UserServiceClient
	utuc *cache.UserToken

	log *log.Helper
}

func NewAuthenticationService(logger log.Logger, sc userv1.UserServiceClient, utuc *cache.UserToken) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "frontend/service/authentication"))
	return &AuthenticationService{
		log:  l,
		sc:   sc,
		utuc: utuc,
	}
}

// Login 登陆
func (s *AuthenticationService) Login(ctx context.Context, req *clientv1.LoginRequest) (*clientv1.LoginResponse, error) {
	if _, err := s.sc.VerifyPassword(ctx, &userv1.VerifyPasswordRequest{
		UserName: req.GetUserName(),
		Password: req.GetPassword(),
	}); err != nil {
		return &clientv1.LoginResponse{}, err
	}

	user, err := s.sc.GetUserByUserName(ctx, &userv1.GetUserByUserNameRequest{UserName: req.GetUserName()})
	if err != nil {
		return &clientv1.LoginResponse{}, err
	}

	accessToken, refreshToken, err := s.utuc.GenerateToken(ctx, user)
	if err != nil {
		return nil, err
	}

	return &clientv1.LoginResponse{
		TokenType:    clientv1.TokenType_bearer.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Logout 登出
func (s *AuthenticationService) Logout(ctx context.Context, req *clientv1.LogoutRequest) (*emptypb.Empty, error) {
	err := s.utuc.RemoveToken(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AuthenticationService) GetMe(ctx context.Context, req *clientv1.GetMeRequest) (*userv1.User, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("User authentication failed [%s]", err.Error())
		return nil, clientv1.ErrorAccessForbidden("User authentication failed")
	}

	req.Id = authInfo.UserId

	return s.sc.GetUser(ctx, &userv1.GetUserRequest{
		Id: req.GetId(),
	})
}
