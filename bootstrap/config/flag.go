package config

import (
	"flag"

	"github.com/realHoangHai/kratos/bootstrap/utils"
)

// CommandFlags Command parameters
type CommandFlags struct {
	Conf       string // Boot configuration file path, default is: ../../configs
	Env        string // Development environment: dev, debug...
	ConfigHost string // Remote configuration server address
	ConfigType string // Remote configuration server type
	Daemon     bool   // Whether to convert to daemon process
}

func NewCommandFlags() *CommandFlags {
	f := &CommandFlags{
		Conf:       "",
		Env:        "",
		ConfigHost: "",
		ConfigType: "",
		Daemon:     false,
	}

	f.defineFlag()

	return f
}

func (f *CommandFlags) defineFlag() {
	flag.StringVar(&f.Conf, "conf", "../../configs", "config path, eg: -conf ../../configs")
	flag.StringVar(&f.Env, "env", "dev", "runtime environment, eg: -env dev")
	flag.StringVar(&f.ConfigHost, "chost", "127.0.0.1:8500", "config server host, eg: -chost 127.0.0.1:8500")
	flag.StringVar(&f.ConfigType, "ctype", "consul", "config server host, eg: -ctype consul")
	flag.BoolVar(&f.Daemon, "d", false, "run app as a daemon with -d=true.")
}

func (f *CommandFlags) Init() {
	flag.Parse()

	if f.Daemon {
		utils.BeDaemon("-d")
	}
}
