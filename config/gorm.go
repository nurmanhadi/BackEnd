package config

import (
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(viper *viper.Viper) *gorm.DB {
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	pool, _ := db.DB()
	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(30)
	pool.SetConnMaxIdleTime(10 * time.Minute)
	pool.SetConnMaxLifetime(30 * time.Minute)
	return db
}
