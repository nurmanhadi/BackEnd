package main

import (
	"fmt"
	"liva/config"
)

func main() {
	viper := config.NewViper()
	app := config.NewFiber(viper)
	db := config.NewConnection(viper)
	log := config.NewLogger()
	validation := config.NewValidator()
	config.New(&config.Bootstrap{
		App:        app,
		DB:         db,
		Log:        log,
		Viper:      viper,
		Validation: validation,
	})

	host := viper.GetString("server.host")
	port := viper.GetInt("server.port")
	if err := app.Listen(fmt.Sprintf("%s:%d", host, port)); err != nil {
		panic(err)
	}
}
