package service

import (
	"context"

	helloworldv1 "github.com/realHoangHai/kratos/api/helloworld/v1"
	"github.com/realHoangHai/kratos/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	helloworldv1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *helloworldv1.HelloRequest) (*helloworldv1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &helloworldv1.HelloReply{Message: "Hello " + g.Hello}, nil
}
