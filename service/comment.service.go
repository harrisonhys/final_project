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

	_comment "final_project/service/comment"
)

type CommentService interface {
	All(userID string) (*[]_comment.CommentResponse, error)
	CreateComment(commentRequest dto.CreateCommentRequest, userID string) (*_comment.CommentResponse, error)
	UpdateComment(updateCommentRequest dto.UpdateCommentRequest, userID string) (*_comment.CommentResponse, error)
	FindOneCommentByID(commentID string) (*_comment.CommentResponse, error)
	DeleteComment(commentID string, userID string) error
}

type commentService struct {
	commentRepo repo.CommentRepository
}

func NewCommentService(commentRepo repo.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

func (c *commentService) All(userID string) (*[]_comment.CommentResponse, error) {
	comments, err := c.commentRepo.All(userID)
	if err != nil {
		return nil, err
	}

	prods := _comment.NewCommentArrayResponse(comments)
	return &prods, nil
}

func (c *commentService) CreateComment(commentRequest dto.CreateCommentRequest, userID string) (*_comment.CommentResponse, error) {
	comment := entity.Comment{}
	err := smapping.FillStruct(&comment, smapping.MapFields(&commentRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	id, _ := strconv.ParseInt(userID, 0, 64)
	comment.UserId = id
	p, err := c.commentRepo.InsertComment(comment)
	if err != nil {
		return nil, err
	}

	res := _comment.NewCommentResponse(p)
	return &res, nil
}

func (c *commentService) FindOneCommentByID(commentID string) (*_comment.CommentResponse, error) {
	comment, err := c.commentRepo.FindOneCommentByID(commentID)

	if err != nil {
		return nil, err
	}

	res := _comment.NewCommentResponse(comment)
	return &res, nil
}

func (c *commentService) UpdateComment(updateCommentRequest dto.UpdateCommentRequest, userID string) (*_comment.CommentResponse, error) {
	comment, err := c.commentRepo.FindOneCommentByID(fmt.Sprintf("%d", updateCommentRequest.ID))
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(userID, 0, 64)
	if comment.UserId != uid {
		return nil, errors.New("Comment ini bukan milik anda")
	}

	comment = entity.Comment{}
	err = smapping.FillStruct(&comment, smapping.MapFields(&updateCommentRequest))

	if err != nil {
		return nil, err
	}

	comment.UserId = uid
	comment, err = c.commentRepo.UpdateComment(comment)

	if err != nil {
		return nil, err
	}

	res := _comment.NewCommentResponse(comment)
	return &res, nil
}

func (c *commentService) DeleteComment(commentID string, userID string) error {
	comment, err := c.commentRepo.FindOneCommentByID(commentID)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%d", comment.UserId) != userID {
		return errors.New("Comment ini bukan milik anda")
	}

	c.commentRepo.DeleteComment(commentID)
	return nil

}
