package rpc

import (
	"net/http/pprof"

	"github.com/gorilla/handlers"

	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/middleware"
	kratosratelimit "github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	kratosrest "github.com/go-kratos/kratos/v2/transport/http"

	"github.com/realHoangHai/kratos/api/gen/go/conf"
)

// CreateHttpServer create http server
func CreateHttpServer(cfg *conf.Bootstrap, m ...middleware.Middleware) *kratosrest.Server {
	var opts []kratosrest.ServerOption

	if cfg.Server.Http.Cors != nil {
		corsOption := kratosrest.Filter(handlers.CORS(
			handlers.AllowedHeaders(cfg.Server.Http.Cors.Headers),
			handlers.AllowedMethods(cfg.Server.Http.Cors.Methods),
			handlers.AllowedOrigins(cfg.Server.Http.Cors.Origins),
		))
		opts = append(opts, corsOption)
	}

	var ms []middleware.Middleware
	if cfg.Server != nil && cfg.Server.Http != nil && cfg.Server.Http.Middleware != nil {
		if cfg.Server.Http.Middleware.GetEnableRecovery() {
			ms = append(ms, recovery.Recovery())
		}
		if cfg.Server.Http.Middleware.GetEnableTracing() {
			ms = append(ms, tracing.Server())
		}
		if cfg.Server.Http.Middleware.GetEnableValidate() {
			ms = append(ms, validate.Validator())
		}
		if cfg.Server.Http.Middleware.GetEnableCircuitBreaker() {
		}
		if cfg.Server.Http.Middleware.Limiter != nil {
			var limiter ratelimit.Limiter
			switch cfg.Server.Http.Middleware.Limiter.GetName() {
			case "bbr":
				limiter = bbr.NewLimiter()
			}
			ms = append(ms, kratosratelimit.Server(kratosratelimit.WithLimiter(limiter)))
		}
	}
	ms = append(ms, m...)
	opts = append(opts, kratosrest.Middleware(ms...))

	if cfg.Server.Http.Network != "" {
		opts = append(opts, kratosrest.Network(cfg.Server.Http.Network))
	}
	if cfg.Server.Http.Addr != "" {
		opts = append(opts, kratosrest.Address(cfg.Server.Http.Addr))
	}
	if cfg.Server.Http.Timeout != nil {
		opts = append(opts, kratosrest.Timeout(cfg.Server.Http.Timeout.AsDuration()))
	}

	srv := kratosrest.NewServer(opts...)

	if cfg.Server.Http.GetEnablePprof() {
		registerHttpPprof(srv)
	}

	return srv
}

func registerHttpPprof(s *kratosrest.Server) {
	s.HandleFunc("/debug/pprof", pprof.Index)
	s.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	s.HandleFunc("/debug/pprof/profile", pprof.Profile)
	s.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	s.HandleFunc("/debug/pprof/trace", pprof.Trace)

	s.HandleFunc("/debug/pprof/allocs", pprof.Handler("allocs").ServeHTTP)
	s.HandleFunc("/debug/pprof/block", pprof.Handler("block").ServeHTTP)
	s.HandleFunc("/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
	s.HandleFunc("/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
	s.HandleFunc("/debug/pprof/mutex", pprof.Handler("mutex").ServeHTTP)
	s.HandleFunc("/debug/pprof/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
}
