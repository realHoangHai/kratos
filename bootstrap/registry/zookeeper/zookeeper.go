package zookeeper

import (
	"github.com/go-zookeeper/zk"

	kratoszookeeper "github.com/go-kratos/kratos/contrib/registry/zookeeper/v2"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/realHoangHai/kratos/api/gen/go/conf"
)

// NewRegistry create new registry - ZooKeeper
func NewRegistry(c *conf.Registry) *kratoszookeeper.Registry {
	conn, _, err := zk.Connect(c.Zookeeper.Endpoints, c.Zookeeper.Timeout.AsDuration())
	if err != nil {
		log.Fatal(err)
	}

	reg := kratoszookeeper.New(conn)
	if err != nil {
		log.Fatal(err)
	}

	return reg
}
