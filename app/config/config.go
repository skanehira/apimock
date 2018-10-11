package config

import (
	"github.com/jinzhu/configor"
	"github.com/skanehira/mockapi/app/common"
)

type Config struct {
	Protocol    string `default:"http"`
	Address     string `default:"localhost"`
	Port        string `default:":8080"`
	CertFile    string `default:"server.crt"`
	CertKeyFile string `default:"server.key"`
	DB          DB
}

type DB struct {
	DBType  string `default:"sqlite3"`
	LogMode bool   `default:"false"`
}

func New(file string) *Config {
	config := new(Config)
	if err := configor.Load(config, file); err != nil {
		panic(common.NewErrLoadConfig(err))
	}

	return config
}
