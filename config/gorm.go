package config

import (
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
	return db
}
