package _comment

import (
	"final_project/entity"
	_user "final_project/service/user"
)

type CommentResponse struct {
	ID      int64              `json:"id"`
	Message string             `json:"message"`
	PhotoId int64              `json:"photo_id"`
	UserId  int64              `json:"user_id"`
	User    _user.UserResponse `json:"user,omitempty"`
}

func NewCommentResponse(comment entity.Comment) CommentResponse {
	return CommentResponse{
		ID:      comment.ID,
		Message: comment.Message,
		PhotoId: comment.PhotoId,
		UserId:  comment.UserId,
		User:    _user.NewUserResponse(comment.User),
	}
}

func NewCommentArrayResponse(comments []entity.Comment) []CommentResponse {
	commentRes := []CommentResponse{}
	for _, v := range comments {
		p := CommentResponse{
			ID:      v.ID,
			Message: v.Message,
			PhotoId: v.PhotoId,
			UserId:  v.UserId,
			User:    _user.NewUserResponse(v.User),
		}
		commentRes = append(commentRes, p)
	}
	return commentRes
}
