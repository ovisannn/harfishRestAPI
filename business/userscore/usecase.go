package userscore

import (
	"context"
	"harfishRestAPI/helpers/messages"
)

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

// result, err := UseCase.categoriesRepo.Create(ctx, data)
// if err != nil {
// 	return Domain{}, err
// }
// return result, nil
func (Usecase *UserScoreUsecase) CreateScoreAndFeedback(ctx context.Context, data *UserScoreDomain) error {
	err := Usecase.UserScoreRepo.CreateScoreAndFeedback(ctx, data)
	if err != nil {
		return err
	}
	return messages.ErrInsertSuccess
}
