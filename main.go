package main

import (
	// "github.com/labstack/echo/middleware"

	"harfishRestAPI/model/mongo"
	"log"

	_routes "harfishRestAPI/app/routes"
	_userscoreUsecase "harfishRestAPI/business/userscore"
	_userscoreController "harfishRestAPI/controllers/userscore"
	_userscoreRepo "harfishRestAPI/model/database"

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

	userscoreRepo := _userscoreRepo.NewMongoUserScore(db)
	userscoreUsecase := _userscoreUsecase.NewUserscoreUsecase(userscoreRepo)
	userscoreController := _userscoreController.NewUserscoreController(userscoreUsecase)

	// init routes

	routesInit := _routes.ControllerList{
		UserControllers: *userscoreController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
