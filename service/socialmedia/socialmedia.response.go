package _socialmedia

import (
	"final_project/entity"
	_user "final_project/service/user"
)

type SocialmediaResponse struct {
	ID             int64              `json:"id"`
	Name           string             `json:"name"`
	SocialMediaUrl string             `json:"social_media_url"`
	User           _user.UserResponse `json:"user,omitempty"`
}

func NewSocialmediaResponse(socialmedia entity.Socialmedia) SocialmediaResponse {
	return SocialmediaResponse{
		ID:             socialmedia.ID,
		Name:           socialmedia.Name,
		SocialMediaUrl: socialmedia.SocialMediaUrl,
		User:           _user.NewUserResponse(socialmedia.User),
	}
}

func NewSocialmediaArrayResponse(socialmedias []entity.Socialmedia) []SocialmediaResponse {
	socialmediaRes := []SocialmediaResponse{}
	for _, v := range socialmedias {
		p := SocialmediaResponse{
			ID:             v.ID,
			Name:           v.Name,
			SocialMediaUrl: v.SocialMediaUrl,
			User:           _user.NewUserResponse(v.User),
		}
		socialmediaRes = append(socialmediaRes, p)
	}
	return socialmediaRes
}
