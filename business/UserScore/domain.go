package userscore

import "context"

type UserScoreDomain struct {
	IdUser   int
	UserName string
	FishName string
	Score    int
	Feedback string
}

type Usecase interface {
}

type Repository interface {
	GetAllScore(ctx context.Context) ([]UserScoreDomain, error)
	// get top 20
}
