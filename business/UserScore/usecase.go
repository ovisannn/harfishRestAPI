package userscore

import "context"

type UserScoreUsecase struct {
	UserScoreRepo Repository
}

func NewUserscoreUsecase(userScoreRepo Repository) Usecase {
	return &UserScoreUsecase{
		UserScoreRepo: userScoreRepo,
	}
}

func (Usecase *UserScoreUsecase) GetAllScore(ctx context.Context) ([]UserScoreDomain, error) {
	result, err := Usecase.UserScoreRepo.GetAllScore(ctx)
	if err != nil {
		return []UserScoreDomain{}, err
	}
	return result, nil
}
