package utility

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type PostgresInfo struct {
	Host                   string `yaml:"host"`
	Port                   string `yaml:"port"`
	Password               string `yaml:"password"`
	Username               string `yaml:"username"`
	Dbname                 string `yaml:"dbname"`
	CloudSqlConnectionName string `yaml:"cloud_sql_connection_name"`
}

type JWTInfo struct {
	Secret string `yaml:"secret"`
}

type ApplicationConfig struct {
	Postgres PostgresInfo `yaml:"postgres"`
	JWT      JWTInfo      `yaml:"jwt"`
}

func LoadApplicationConfig(configDir, configFile string) (*ApplicationConfig, error) {
	content, err := ioutil.ReadFile((filepath.Join(configDir, configFile)))
	if err != nil {
		return nil, err
	}

	var config ApplicationConfig
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
