package repo

import (
	"final_project/entity"

	"gorm.io/gorm"
)

type CommentRepository interface {
	All(userID string) ([]entity.Comment, error)
	InsertComment(comment entity.Comment) (entity.Comment, error)
	UpdateComment(comment entity.Comment) (entity.Comment, error)
	DeleteComment(commentID string) error
	FindOneCommentByID(ID string) (entity.Comment, error)
	FindAllComment(userID string) ([]entity.Comment, error)
}

type commentRepo struct {
	connection *gorm.DB
}

func NewCommentRepo(connection *gorm.DB) CommentRepository {
	return &commentRepo{
		connection: connection,
	}
}

func (c *commentRepo) All(userID string) ([]entity.Comment, error) {
	comments := []entity.Comment{}
	c.connection.Preload("User").Where("user_id = ?", userID).Find(&comments)
	return comments, nil
}

func (c *commentRepo) InsertComment(comment entity.Comment) (entity.Comment, error) {
	c.connection.Save(&comment)
	c.connection.Preload("User").Find(&comment)
	return comment, nil
}

func (c *commentRepo) UpdateComment(comment entity.Comment) (entity.Comment, error) {
	c.connection.Save(&comment)
	c.connection.Preload("User").Find(&comment)
	return comment, nil
}

func (c *commentRepo) FindOneCommentByID(commentID string) (entity.Comment, error) {
	var comment entity.Comment
	res := c.connection.Preload("User").Where("id = ?", commentID).Take(&comment)
	if res.Error != nil {
		return comment, res.Error
	}
	return comment, nil
}

func (c *commentRepo) FindAllComment(userID string) ([]entity.Comment, error) {
	comments := []entity.Comment{}
	c.connection.Where("user_id = ?", userID).Find(&comments)
	return comments, nil
}

func (c *commentRepo) DeleteComment(commentID string) error {
	var comment entity.Comment
	res := c.connection.Preload("User").Where("id = ?", commentID).Take(&comment)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&comment)
	return nil
}
