package consul

import (
	"strings"

	consulapi "github.com/hashicorp/consul/api"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/realHoangHai/kratos/api/gen/go/conf"
)

// getConfigKey get the legal configuration name
func getConfigKey(configKey string, useBackslash bool) string {
	if useBackslash {
		return strings.Replace(configKey, `.`, `/`, -1)
	} else {
		return configKey
	}
}

// NewConfigSource create a remote configuration source - Consul
func NewConfigSource(c *conf.RemoteConfig) config.Source {
	cfg := consulapi.DefaultConfig()
	cfg.Address = c.Consul.Address
	cfg.Scheme = c.Consul.Scheme

	cli, err := consulapi.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	src, err := New(cli,
		WithPath(getConfigKey(c.Consul.Key, true)),
	)
	if err != nil {
		log.Fatal(err)
	}

	return src
}
