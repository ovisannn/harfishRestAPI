package userscore

import "context"

type UserScoreDomain struct {
	// IdUser   int
	UserName string `bson:"username,omitempty" json:"username,omitempty"`
	FishName string `bson:"fishname,omitempty" json:"fishname,omitempty"`
	Score    int    `bson:"score,omitempty" json:"score,omitempty"`
	Feedback string `bson:"feedback,omitempty" json:"feedback,omitempty"`
}

type Usecase interface {
	GetAllScore(ctx context.Context) ([]UserScoreDomain, error)
}

type Repository interface {
	GetAllScore(ctx context.Context) ([]UserScoreDomain, error)
	// get top 20
}
