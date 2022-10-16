package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"final_project/common/obj"
	"final_project/common/response"
	"final_project/dto"
	"final_project/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CommentHandler interface {
	All(ctx *gin.Context)
	CreateComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
	FindOneCommentByID(ctx *gin.Context)
}

type commentHandler struct {
	commentService service.CommentService
	jwtService     service.JWTService
}

func NewCommentHandler(commentService service.CommentService, jwtService service.JWTService) CommentHandler {
	return &commentHandler{
		commentService: commentService,
		jwtService:     jwtService,
	}
}

func (c *commentHandler) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	comments, err := c.commentService.All(userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", comments)
	ctx.JSON(http.StatusOK, response)
}

func (c *commentHandler) CreateComment(ctx *gin.Context) {
	var createCommentReq dto.CreateCommentRequest
	err := ctx.ShouldBind(&createCommentReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.commentService.CreateComment(createCommentReq, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *commentHandler) FindOneCommentByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.commentService.FindOneCommentByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *commentHandler) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	err := c.commentService.DeleteComment(id, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *commentHandler) UpdateComment(ctx *gin.Context) {
	updateCommentRequest := dto.UpdateCommentRequest{}
	err := ctx.ShouldBind(&updateCommentRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateCommentRequest.ID = id
	comment, err := c.commentService.UpdateComment(updateCommentRequest, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", comment)
	ctx.JSON(http.StatusOK, response)

}
