package registry

import (
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/realHoangHai/kratos/bootstrap/registry/zookeeper"
	"github.com/realHoangHai/kratos/conf"
)

type Type string

const (
	Consul     Type = "consul"
	Etcd       Type = "etcd"
	ZooKeeper  Type = "zookeeper"
	Nacos      Type = "nacos"
	Kubernetes Type = "kubernetes"
)

// NewRegistry create a new registry
func NewRegistry(cfg *conf.Registry) registry.Registrar {
	if cfg == nil {
		return nil
	}

	switch Type(cfg.Type) {
	case ZooKeeper:
		return zookeeper.NewRegistry(cfg)
	}

	return nil
}

// NewDiscovery create new discovery
func NewDiscovery(cfg *conf.Registry) registry.Discovery {
	if cfg == nil {
		return nil
	}

	switch Type(cfg.Type) {
	case ZooKeeper:
		return zookeeper.NewRegistry(cfg)
	}

	return nil
}
