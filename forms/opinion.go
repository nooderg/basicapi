package forms

import "basic-api/models"

type OpignionForm struct {
	Like bool `json:"like"`
}

func (o OpignionForm) PrepareOpinion(userID uint, articleID uint) models.Opinion {
	return models.Opinion{
		UserID:    uint(userID),
		ArticleID: articleID,
		Like:      o.Like,
	}
}
