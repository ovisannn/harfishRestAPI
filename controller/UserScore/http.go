package userscore

import userscore "harfishRestAPI/business/UserScore"

type UserscoreController struct {
	UserScoreUsecase userscore.Usecase
}

func NewUserscoreController(userscoreUsecase userscore.Usecase) *UserscoreController {
	return &UserscoreController{
		UserScoreUsecase: userscoreUsecase,
	}
}
