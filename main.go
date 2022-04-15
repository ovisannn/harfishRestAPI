package main

import (
	// "github.com/labstack/echo/middleware"

	"harfishRestAPI/model/mongo"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("app/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

}

func main() {
	config := mongo.Config{
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Name:     viper.GetString("database.name"),
	}
	e := echo.New()
	CorsHandle := viper.GetStringSlice("CORS.AllowOrigins")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: CorsHandle,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	db, errDb := config.ConnectDB()
	if errDb != nil {
		panic(errDb)
	}

	// init repo
	// init usecase
	// init controller

	// init routes

}
