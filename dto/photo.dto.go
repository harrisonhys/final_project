package dto

type CreatePhotoRequest struct {
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" form:"photo_url" binding:"required"`
}

type UpdatePhotoRequest struct {
	ID       int64  `json:"id" form:"id"`
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" form:"photo_url" binding:"required"`
}
