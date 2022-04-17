package routes

import (
	_userscoreController "harfishRestAPI/controllers/userscore"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserscoreController _userscoreController.UserscoreController
}

func (ctrl *ControllerList) RouteRegister(e *echo.Echo) {
	baseRoute := e.Group("/harfish")
	baseRoute.GET("/userscore", ctrl.UserscoreController.GetAll)
	baseRoute.POST("/userscore", ctrl.UserscoreController.CreateScoreAndFeedback)
	//post score
	//get best 10
}
