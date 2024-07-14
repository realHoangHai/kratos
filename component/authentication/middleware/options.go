package middleware

import (
	"context"
	"github.com/realHoangHai/kratos/component/authentication/engine"
)

type Option func(*options)

type options struct {
	claims engine.AuthClaims
}

func WithAuthClaims(claims engine.AuthClaims) Option {
	return func(o *options) {
		o.claims = claims
	}
}

func NewContext(ctx context.Context, claims *engine.AuthClaims) context.Context {
	return engine.ContextWithAuthClaims(ctx, claims)
}

func FromContext(ctx context.Context) (*engine.AuthClaims, bool) {
	return engine.AuthClaimsFromContext(ctx)
}
