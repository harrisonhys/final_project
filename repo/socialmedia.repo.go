package repo

import (
	"final_project/entity"

	"gorm.io/gorm"
)

type SocialmediaRepository interface {
	All(userID string) ([]entity.Socialmedia, error)
	InsertSocialmedia(socialmedia entity.Socialmedia) (entity.Socialmedia, error)
	UpdateSocialmedia(socialmedia entity.Socialmedia) (entity.Socialmedia, error)
	DeleteSocialmedia(socialmediaID string) error
	FindOneSocialmediaByID(ID string) (entity.Socialmedia, error)
	FindAllSocialmedia(userID string) ([]entity.Socialmedia, error)
}

type socialmediaRepo struct {
	connection *gorm.DB
}

func NewSocialmediaRepo(connection *gorm.DB) SocialmediaRepository {
	return &socialmediaRepo{
		connection: connection,
	}
}

func (c *socialmediaRepo) All(userID string) ([]entity.Socialmedia, error) {
	socialmedias := []entity.Socialmedia{}
	c.connection.Preload("User").Where("user_id = ?", userID).Find(&socialmedias)
	return socialmedias, nil
}

func (c *socialmediaRepo) InsertSocialmedia(socialmedia entity.Socialmedia) (entity.Socialmedia, error) {
	c.connection.Save(&socialmedia)
	c.connection.Preload("User").Find(&socialmedia)
	return socialmedia, nil
}

func (c *socialmediaRepo) UpdateSocialmedia(socialmedia entity.Socialmedia) (entity.Socialmedia, error) {
	c.connection.Save(&socialmedia)
	c.connection.Preload("User").Find(&socialmedia)
	return socialmedia, nil
}

func (c *socialmediaRepo) FindOneSocialmediaByID(socialmediaID string) (entity.Socialmedia, error) {
	var socialmedia entity.Socialmedia
	res := c.connection.Preload("User").Where("id = ?", socialmediaID).Take(&socialmedia)
	if res.Error != nil {
		return socialmedia, res.Error
	}
	return socialmedia, nil
}

func (c *socialmediaRepo) FindAllSocialmedia(userID string) ([]entity.Socialmedia, error) {
	socialmedias := []entity.Socialmedia{}
	c.connection.Where("user_id = ?", userID).Find(&socialmedias)
	return socialmedias, nil
}

func (c *socialmediaRepo) DeleteSocialmedia(socialmediaID string) error {
	var socialmedia entity.Socialmedia
	res := c.connection.Preload("User").Where("id = ?", socialmediaID).Take(&socialmedia)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&socialmedia)
	return nil
}
