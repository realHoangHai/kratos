package service

import (
	"github.com/go-kratos/kratos/v2/log"

	permissionv1 "github.com/realHoangHai/kratos/api/gen/go/permission/v1"
	"github.com/realHoangHai/kratos/app/core/internal/data"
)

type PermissionService struct {
	permissionv1.UnimplementedPermissionServiceServer

	uc  *data.PermissionRepo
	log *log.Helper
}

func NewPermissionService(logger log.Logger, uc *data.PermissionRepo) *PermissionService {
	l := log.NewHelper(log.With(logger, "module", "permission/service/core-service"))
	return &PermissionService{
		uc:  uc,
		log: l,
	}
}
