package database

import userscore "harfishRestAPI/business/UserScore"

type UserScore struct {
	IdUser   int
	UserName string
	FishName string
	Score    int
	Feedback string
}

func (record *UserScore) UsesrScoreToDomain() userscore.UserScoreDomain {
	return userscore.UserScoreDomain{
		IdUser:   record.IdUser,
		UserName: record.UserName,
		FishName: record.FishName,
		Score:    record.Score,
		Feedback: record.Feedback,
	}
}

func UserScoreFromDomain(domain userscore.UserScoreDomain) UserScore {
	return UserScore{
		IdUser:   domain.IdUser,
		UserName: domain.UserName,
		FishName: domain.FishName,
		Score:    domain.Score,
		Feedback: domain.Feedback,
	}
}
