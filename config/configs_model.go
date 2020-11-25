package configs

import (
	o_mgo "git.52retail.com/oliver/oliver/db/o.mgo"
)

// Config is the instance of configuration load from config file
var Config Configuration

// Configuration 模块配置
type Configuration struct {
	Stage           string            `json:"stage"`
	BasicAuth       *BasicAuth        `json:"basicAuth"`
	EcnsMap         map[string]string `json:"ecnsMap"`
	ValueMap        map[string]int    `json:"valueMap"`
	GcnsMap         map[string]string `json:"gcnsMap"`
	PushDate        string            `json:"pushDate"`
	TplKey          string            `json:"tplKey"`
	EmailNotifyList []string
	PostgresConf    *DatabaseConfig
	MongoConf       *o_mgo.MgoConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

type BasicAuth struct {
	Account  string
	Password string
}
