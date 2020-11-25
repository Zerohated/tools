package dao

import (
	o_mgo "git.52retail.com/oliver/oliver/db/o.mgo"
)

var (
	MongoManager *o_mgo.MongoSessionManager
)

func init() {
	mgoConfig := &o_mgo.MgoConfig{
		ModuleName:    config.MongoConf.ModuleName,
		ConnectionStr: config.MongoConf.ConnectionStr,
		PoolLimit:     config.MongoConf.PoolLimit,
	}
	MongoManager = o_mgo.NewMongoSessionManager(mgoConfig)
}

// CloseDB 关闭数据库连接
func CloseDB() {
	if MongoManager != nil {
		MongoManager.Dispose()
	}
}
