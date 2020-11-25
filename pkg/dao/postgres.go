package dao

import (
	"fmt"
	"time"

	conf "github.com/Zerohated/tools/config"

	"github.com/jinzhu/gorm"
	// import postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	config = conf.Config
	// PgConn it is safe for goroutines. Use PgConn for all postgres request.
	PgConn = connectPG()
)

func connectPG() *gorm.DB {
	pgConf := config.PostgresConf
	dbEndPoint := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=2", pgConf.Host, pgConf.Port, pgConf.User, pgConf.Password, pgConf.DBName)
	PgConn, err := gorm.Open("postgres", dbEndPoint)
	if err != nil {
		fmt.Println(err.Error())
	}

	PgConn.DB().SetConnMaxLifetime(time.Minute * 5)
	PgConn.DB().SetMaxIdleConns(20)
	PgConn.DB().SetMaxOpenConns(500)
	return PgConn
}
