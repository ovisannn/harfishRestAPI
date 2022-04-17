package response

import userscoreU "harfishRestAPI/business/userscore"

type UserScore struct {
	// IdUser   int
	UserName string `bson:"username,omitempty" json:"username,omitempty"`
	FishName string `bson:"fishname,omitempty" json:"fishname,omitempty"`
	Score    int    `bson:"score,omitempty" json:"score,omitempty"`
	Feedback string `bson:"feedback,omitempty" json:"feedback,omitempty"`
}

func UserScoreFromDomain(domain userscoreU.UserScoreDomain) UserScore {
	return UserScore{
		// IdUser:   domain.IdUser,
		UserName: domain.UserName,
		FishName: domain.FishName,
		Score:    domain.Score,
		Feedback: domain.Feedback,
	}
}
