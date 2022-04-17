package userscore

import (
	"context"
	"harfishRestAPI/helpers/messages"
	"sort"
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
	sort.Slice(result, func(i, j int) bool {
		return result[i].Score > result[j].Score
	})
	return result, nil
}

func (Usecase *UserScoreUsecase) CreateScoreAndFeedback(ctx context.Context, data *UserScoreDomain) error {
	err := Usecase.UserScoreRepo.CreateScoreAndFeedback(ctx, data)
	if err != nil {
		return err
	}
	return messages.ErrInsertSuccess
}
