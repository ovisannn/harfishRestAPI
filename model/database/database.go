package database

import userscoreU "harfishRestAPI/business/userscore"

type UserScore struct {
	IdUser   int
	UserName string
	FishName string
	Score    int
	Feedback string
}

func (record *UserScore) UsesrScoreToDomain() userscoreU.UserScoreDomain {
	return userscoreU.UserScoreDomain{
		IdUser:   record.IdUser,
		UserName: record.UserName,
		FishName: record.FishName,
		Score:    record.Score,
		Feedback: record.Feedback,
	}
}

func UserScoreFromDomain(domain userscoreU.UserScoreDomain) UserScore {
	return UserScore{
		IdUser:   domain.IdUser,
		UserName: domain.UserName,
		FishName: domain.FishName,
		Score:    domain.Score,
		Feedback: domain.Feedback,
	}
}
