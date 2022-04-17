package response

import userscoreU "harfishRestAPI/business/userscore"

type UserScore struct {
	IdUser   int
	UserName string
	FishName string
	Score    int
	Feedback string
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
