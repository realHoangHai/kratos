package handler

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/go-kratos/kratos/v2/log"

	userv1 "github.com/realHoangHai/kratos/api/gen/go/user/v1"
	authn "github.com/realHoangHai/kratos/pkg/authentication/engine"
)

type UserToken struct {
	log           *log.Helper
	rdb           *redis.Client
	authenticator authn.Authenticator

	accessTokenKeyPrefix  string
	refreshTokenKeyPrefix string
}

func NewUserToken(
	rdb *redis.Client,
	authenticator authn.Authenticator,
	logger log.Logger,
	accessTokenKeyPrefix string,
	refreshTokenKeyPrefix string,
) *UserToken {
	l := log.NewHelper(log.With(logger, "module", "handler/user-token"))
	return &UserToken{
		log:                   l,
		rdb:                   rdb,
		authenticator:         authenticator,
		accessTokenKeyPrefix:  accessTokenKeyPrefix,
		refreshTokenKeyPrefix: refreshTokenKeyPrefix,
	}
}

// createAccessJwtToken generate JWT access token
func (r *UserToken) createAccessJwtToken(_ string, userId uint32) string {
	principal := authn.AuthClaims{
		Subject: strconv.FormatUint(uint64(userId), 10),
		Scopes:  make(authn.ScopeSet),
	}

	signedToken, err := r.authenticator.CreateIdentity(principal)
	if err != nil {
		return ""
	}

	return signedToken
}

// createRefreshToken generate refresh token
func (r *UserToken) createRefreshToken() string {
	return uuid.New().String()
}

// GenerateToken create token
func (r *UserToken) GenerateToken(ctx context.Context, user *userv1.User) (accessToken string, refreshToken string, err error) {
	if accessToken = r.createAccessJwtToken(user.GetUserName(), user.GetId()); accessToken == "" {
		err = errors.New("create access token failed")
		return
	}

	if err = r.setAccessTokenToRedis(ctx, user.GetId(), accessToken, 0); err != nil {
		return
	}

	if refreshToken = r.createRefreshToken(); refreshToken == "" {
		err = errors.New("create refresh token failed")
		return
	}

	if err = r.setRefreshTokenToRedis(ctx, user.GetId(), refreshToken, 0); err != nil {
		return
	}

	return
}

// GenerateAccessToken create access token
func (r *UserToken) GenerateAccessToken(ctx context.Context, userId uint32, userName string) (accessToken string, err error) {
	if accessToken = r.createAccessJwtToken(userName, userId); accessToken == "" {
		err = errors.New("create access token failed")
		return
	}

	if err = r.setAccessTokenToRedis(ctx, userId, accessToken, 0); err != nil {
		return
	}

	return
}

// GenerateRefreshToken create refresh token
func (r *UserToken) GenerateRefreshToken(ctx context.Context, user *userv1.User) (refreshToken string, err error) {
	if refreshToken = r.createRefreshToken(); refreshToken == "" {
		err = errors.New("create refresh token failed")
		return
	}

	if err = r.setRefreshTokenToRedis(ctx, user.GetId(), refreshToken, 0); err != nil {
		return
	}

	return
}

// RemoveToken remove all tokens
func (r *UserToken) RemoveToken(ctx context.Context, userId uint32) error {
	var err error
	if err = r.deleteAccessTokenFromRedis(ctx, userId); err != nil {
		r.log.Errorf("Remove user access token failed: [%v]", err)
	}

	if err = r.deleteRefreshTokenFromRedis(ctx, userId); err != nil {
		r.log.Errorf("Remove user refresh token failed: [%v]", err)
	}

	return err
}

// GetAccessToken get access token
func (r *UserToken) GetAccessToken(ctx context.Context, userId uint32) string {
	return r.getAccessTokenFromRedis(ctx, userId)
}

// GetRefreshToken get refresh token
func (r *UserToken) GetRefreshToken(ctx context.Context, userId uint32) string {
	return r.getRefreshTokenFromRedis(ctx, userId)
}

func (r *UserToken) setAccessTokenToRedis(ctx context.Context, userId uint32, token string, expires int32) error {
	key := fmt.Sprintf("%s%d", r.accessTokenKeyPrefix, userId)
	return r.rdb.Set(ctx, key, token, time.Duration(expires)).Err()
}

func (r *UserToken) getAccessTokenFromRedis(ctx context.Context, userId uint32) string {
	key := fmt.Sprintf("%s%d", r.accessTokenKeyPrefix, userId)
	result, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("Get redis user access token failed: %s", err.Error())
		}
		return ""
	}
	return result
}

func (r *UserToken) deleteAccessTokenFromRedis(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", r.accessTokenKeyPrefix, userId)
	return r.rdb.Del(ctx, key).Err()
}

func (r *UserToken) setRefreshTokenToRedis(ctx context.Context, userId uint32, token string, expires int32) error {
	key := fmt.Sprintf("%s%d", r.refreshTokenKeyPrefix, userId)
	return r.rdb.Set(ctx, key, token, time.Duration(expires)).Err()
}

func (r *UserToken) getRefreshTokenFromRedis(ctx context.Context, userId uint32) string {
	key := fmt.Sprintf("%s%d", r.refreshTokenKeyPrefix, userId)
	result, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("Get redis user refresh token failed: %s", err.Error())
		}
		return ""
	}
	return result
}

func (r *UserToken) deleteRefreshTokenFromRedis(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", r.refreshTokenKeyPrefix, userId)
	return r.rdb.Del(ctx, key).Err()
}
