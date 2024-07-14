package config

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/realHoangHai/kratos/api/gen/go/conf"
	"github.com/realHoangHai/kratos/bootstrap/config/consul"
	"github.com/realHoangHai/kratos/pkg/utils"
)

// NewRemoteConfigSource create a remote configuration source
func NewRemoteConfigSource(c *conf.RemoteConfig) config.Source {
	switch Type(c.Type) {
	default:
		fallthrough
	case LocalFile:
		return nil
	case Nacos:
		return nil
	case Consul:
		return consul.NewConfigSource(c)
	case Etcd:
		return nil
	case Apollo:
		return nil
	case Kubernetes:
		return nil
	case Polaris:
		return nil
	}
}

// NewFileConfigSource create a local file configuration source
func NewFileConfigSource(filePath string) config.Source {
	return file.NewSource(filePath)
}

// NewConfigProvider create a configuration provider
func NewConfigProvider(configPath string) config.Config {
	err, rc := LoadRemoteConfigSourceConfigs(configPath)
	if err != nil {
		log.Error("LoadRemoteConfigSourceConfigs: ", err.Error())
	}
	if rc != nil {
		return config.New(
			config.WithSource(
				NewFileConfigSource(configPath),
				NewRemoteConfigSource(rc),
			),
		)
	} else {
		return config.New(
			config.WithSource(
				NewFileConfigSource(configPath),
			),
		)
	}
}

// LoadBootstrapConfig loader boot configuration
func LoadBootstrapConfig(configPath string) error {
	cfg := NewConfigProvider(configPath)

	var err error

	if err = cfg.Load(); err != nil {
		return err
	}

	initBootstrapConfig()

	if err = scanConfigs(cfg); err != nil {
		return err
	}

	return nil
}

func scanConfigs(cfg config.Config) error {
	initBootstrapConfig()

	for _, c := range configList {
		if err := cfg.Scan(c); err != nil {
			return err
		}
	}
	return nil
}

// LoadRemoteConfigSourceConfigs load local configuration from remote configuration source
func LoadRemoteConfigSourceConfigs(configPath string) (error, *conf.RemoteConfig) {
	configPath = configPath + "/" + remoteConfigSourceConfigFile
	if !utils.FileExists(configPath) {
		return nil, nil
	}

	cfg := config.New(
		config.WithSource(
			NewFileConfigSource(configPath),
		),
	)
	defer func(cfg config.Config) {
		if err := cfg.Close(); err != nil {
			panic(err)
		}
	}(cfg)

	var err error

	if err = cfg.Load(); err != nil {
		return err, nil
	}

	if err = scanConfigs(cfg); err != nil {
		return err, nil
	}

	return nil, GetBootstrapConfig().Config
}
