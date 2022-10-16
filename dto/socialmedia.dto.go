package dto

type CreateSocialMediaRequest struct {
	Name           string `json:"name" form:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" binding:"required"`
	// UserId         int64  `json:"userid" form:"userid" binding:"required"`
}

type UpdateSocialMediaRequest struct {
	ID             int64  `json:"id" form:"id"`
	Name           string `json:"name" form:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" binding:"required"`
	// UserId         int64  `json:"userid" form:"userid" binding:"required"`
}
