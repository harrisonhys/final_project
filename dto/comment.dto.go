package dto

type CreateCommentRequest struct {
	Message string `json:"message" form:"message" binding:"required"`
	PhotoId int64  `json:"photo_id" form:"photo_id" binding:"required"`
}

type UpdateCommentRequest struct {
	ID      int64  `json:"id" form:"id"`
	Message string `json:"message" form:"message" binding:"required"`
	PhotoId int64  `json:"photo_id" form:"photo_id" binding:"required"`
}
