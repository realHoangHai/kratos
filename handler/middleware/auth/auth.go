package auth

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"strconv"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	authnengine "github.com/realHoangHai/kratos/component/authentication/engine"
	authn "github.com/realHoangHai/kratos/component/authentication/middleware"

	authzengine "github.com/realHoangHai/kratos/component/authorization/engine"
	authz "github.com/realHoangHai/kratos/component/authorization/middleware"
)

const (
	reason string = "UNAUTHORIZED"
)

var (
	ErrWrongContext         = errors.Unauthorized(reason, "wrong context for middleware")
	ErrMissingJwtToken      = errors.Unauthorized(reason, "no jwt token in context")
	ErrExtractSubjectFailed = errors.Unauthorized(reason, "extract subject failed")
)

var action = authzengine.Action("ANY")

// Server connecting authentication and authorization
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			authnClaims, ok := authn.FromContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			sub := authzengine.Subject(authnClaims.Subject)
			path := authzengine.Resource(tr.Operation())
			authzClaims := authzengine.AuthClaims{
				Subject:  &sub,
				Action:   &action,
				Resource: &path,
			}

			ctx = authz.NewContext(ctx, &authzClaims)

			return handler(ctx, req)
		}
	}
}

type Result struct {
	UserId   uint32
	UserName string
}

func FromContext(ctx context.Context) (*Result, error) {
	claims, ok := authnengine.AuthClaimsFromContext(ctx)
	if !ok {
		return nil, ErrMissingJwtToken
	}

	userId, err := strconv.ParseUint(claims.Subject, 10, 32)
	if err != nil {
		return nil, ErrExtractSubjectFailed
	}

	return &Result{
		UserId:   uint32(userId),
		UserName: "",
	}, nil
}
