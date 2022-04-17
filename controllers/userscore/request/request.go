package request

import userscoreU "harfishRestAPI/business/userscore"

type UserScore struct {
	// IdUser   int
	UserName string `bson:"username,omitempty" json:"username,omitempty"`
	FishName string `bson:"fishname,omitempty" json:"fishname,omitempty"`
	Score    int    `bson:"score,omitempty" json:"score,omitempty"`
	Feedback string `bson:"feedback,omitempty" json:"feedback,omitempty"`
}

func (record *UserScore) UsesrScoreToDomain() *userscoreU.UserScoreDomain {
	return &userscoreU.UserScoreDomain{
		// IdUser:   record.IdUser,
		UserName: record.UserName,
		FishName: record.FishName,
		Score:    record.Score,
		Feedback: record.Feedback,
	}
}

// func (record *Categories) ToDomain() *categories.Domain {
// 	return &categories.Domain{
// 		ID:           record.ID,
// 		CategoryName: record.CategoryName,
// 		Description:  record.Description,
// 		Rules:        record.Rules,
// 		Header:       record.Header,
// 		ColorTheme:   record.ColorTheme,
// 	}
// }
