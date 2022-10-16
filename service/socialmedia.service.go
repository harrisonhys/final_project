package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"final_project/dto"
	"final_project/entity"
	"final_project/repo"

	"github.com/mashingan/smapping"

	_socialmedia "final_project/service/socialmedia"
)

type SocialMediaService interface {
	All(userID string) (*[]_socialmedia.SocialmediaResponse, error)
	CreateSocialMedia(socialmediaRequest dto.CreateSocialMediaRequest, userID string) (*_socialmedia.SocialmediaResponse, error)
	UpdateSocialMedia(updateSocialMediaRequest dto.UpdateSocialMediaRequest, userID string) (*_socialmedia.SocialmediaResponse, error)
	FindOneSocialMediaByID(socialmediaID string) (*_socialmedia.SocialmediaResponse, error)
	DeleteSocialMedia(socialmediaID string, userID string) error
}

type socialmediaService struct {
	socialmediaRepo repo.SocialmediaRepository
}

func NewSocialMediaService(socialmediaRepo repo.SocialmediaRepository) SocialMediaService {
	return &socialmediaService{
		socialmediaRepo: socialmediaRepo,
	}
}

func (c *socialmediaService) All(userID string) (*[]_socialmedia.SocialmediaResponse, error) {
	socialmedias, err := c.socialmediaRepo.All(userID)
	if err != nil {
		return nil, err
	}

	prods := _socialmedia.NewSocialmediaArrayResponse(socialmedias)
	return &prods, nil
}

func (c *socialmediaService) CreateSocialMedia(socialmediaRequest dto.CreateSocialMediaRequest, userID string) (*_socialmedia.SocialmediaResponse, error) {
	socialmedia := entity.Socialmedia{}
	err := smapping.FillStruct(&socialmedia, smapping.MapFields(&socialmediaRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	id, _ := strconv.ParseInt(userID, 0, 64)
	socialmedia.Userid = id
	p, err := c.socialmediaRepo.InsertSocialmedia(socialmedia)
	if err != nil {
		return nil, err
	}

	res := _socialmedia.NewSocialmediaResponse(p)
	return &res, nil
}

func (c *socialmediaService) FindOneSocialMediaByID(socialmediaID string) (*_socialmedia.SocialmediaResponse, error) {
	socialmedia, err := c.socialmediaRepo.FindOneSocialmediaByID(socialmediaID)

	if err != nil {
		return nil, err
	}

	res := _socialmedia.NewSocialmediaResponse(socialmedia)
	return &res, nil
}

func (c *socialmediaService) UpdateSocialMedia(updateSocialMediaRequest dto.UpdateSocialMediaRequest, userID string) (*_socialmedia.SocialmediaResponse, error) {
	socialmedia, err := c.socialmediaRepo.FindOneSocialmediaByID(fmt.Sprintf("%d", updateSocialMediaRequest.ID))
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(userID, 0, 64)
	if socialmedia.Userid != uid {
		return nil, errors.New("Sosmed ini bukan milik anda")
	}

	socialmedia = entity.Socialmedia{}
	err = smapping.FillStruct(&socialmedia, smapping.MapFields(&updateSocialMediaRequest))

	if err != nil {
		return nil, err
	}

	socialmedia.Userid = uid
	socialmedia, err = c.socialmediaRepo.UpdateSocialmedia(socialmedia)

	if err != nil {
		return nil, err
	}

	res := _socialmedia.NewSocialmediaResponse(socialmedia)
	return &res, nil
}

func (c *socialmediaService) DeleteSocialMedia(socialmediaID string, userID string) error {
	socialmedia, err := c.socialmediaRepo.FindOneSocialmediaByID(socialmediaID)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%d", socialmedia.Userid) != userID {
		return errors.New("Sosmed ini bukan milik anda")
	}

	c.socialmediaRepo.DeleteSocialmedia(socialmediaID)
	return nil

}
