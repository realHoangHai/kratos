package data

import "github.com/go-kratos/kratos/v2/log"

type PermissionRepo struct {
	data *Data
	log  *log.Helper
}

func NewPermissionRepo(data *Data, logger log.Logger) *PermissionRepo {
	l := log.NewHelper(log.With(logger, "module", "permission/data/core-service"))
	return &PermissionRepo{
		data: data,
		log:  l,
	}
}
