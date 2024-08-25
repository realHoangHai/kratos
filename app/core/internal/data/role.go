package data

import "github.com/go-kratos/kratos/v2/log"

type RoleRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleRepo(data *Data, logger log.Logger) *RoleRepo {
	l := log.NewHelper(log.With(logger, "module", "role/data/core-service"))
	return &RoleRepo{
		data: data,
		log:  l,
	}
}
