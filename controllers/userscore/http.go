package userscore

import (
	userscoreU "harfishRestAPI/business/userscore"
	"harfishRestAPI/controllers"
	"harfishRestAPI/controllers/userscore/request"
	"harfishRestAPI/controllers/userscore/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserscoreController struct {
	UserScoreUsecase userscoreU.Usecase
}

func NewUserscoreController(userscoreUsecase userscoreU.Usecase) *UserscoreController {
	return &UserscoreController{
		UserScoreUsecase: userscoreUsecase,
	}
}

func (controller *UserscoreController) GetAll(c echo.Context) error {
	UserscoreResult := []response.UserScore{}
	ctx := c.Request().Context()

	result, err := controller.UserScoreUsecase.GetAllScore(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	for _, i := range result {
		UserscoreResult = append(UserscoreResult, response.UserScoreFromDomain(i))
	}
	return controllers.NewSuccessResponse(c, UserscoreResult)
}

func (controller *UserscoreController) CreateScoreAndFeedback(c echo.Context) error {
	insertScoreFeedback := request.UserScore{}
	c.Bind(&insertScoreFeedback)
	ctx := c.Request().Context()
	err := controller.UserScoreUsecase.CreateScoreAndFeedback(ctx, insertScoreFeedback.UsesrScoreToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, err)
}
