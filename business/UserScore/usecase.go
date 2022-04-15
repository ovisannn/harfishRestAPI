package userscore

type UserScoreUsecase struct {
	UserScoreRepo Repository
}

func NewUserscoreUsecase(userScoreRepo Repository) Usecase {
	return &UserScoreUsecase{
		UserScoreRepo: userScoreRepo,
	}
}
