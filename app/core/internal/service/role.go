package service

import (
	"github.com/go-kratos/kratos/v2/log"

	rolev1 "github.com/realHoangHai/kratos/api/gen/go/role/v1"
	"github.com/realHoangHai/kratos/app/core/internal/data"
)

type RoleService struct {
	rolev1.UnimplementedRoleServiceServer

	uc  *data.RoleRepo
	log *log.Helper
}

func NewRoleService(logger log.Logger, uc *data.RoleRepo) *RoleService {
	l := log.NewHelper(log.With(logger, "module", "role/service/core-service"))
	return &RoleService{
		uc:  uc,
		log: l,
	}
}
