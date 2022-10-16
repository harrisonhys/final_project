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

type PhotoHandler interface {
	All(ctx *gin.Context)
	CreatePhoto(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
	FindOnePhotoByID(ctx *gin.Context)
}

type photoHandler struct {
	photoService service.PhotoService
	jwtService   service.JWTService
}

func NewPhotoHandler(photoService service.PhotoService, jwtService service.JWTService) PhotoHandler {
	return &photoHandler{
		photoService: photoService,
		jwtService:   jwtService,
	}
}

func (c *photoHandler) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	photos, err := c.photoService.All(userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", photos)
	ctx.JSON(http.StatusOK, response)
}

func (c *photoHandler) CreatePhoto(ctx *gin.Context) {
	var createPhotoReq dto.CreatePhotoRequest
	err := ctx.ShouldBind(&createPhotoReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.photoService.CreatePhoto(createPhotoReq, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *photoHandler) FindOnePhotoByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.photoService.FindOnePhotoByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *photoHandler) DeletePhoto(ctx *gin.Context) {
	id := ctx.Param("id")

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	err := c.photoService.DeletePhoto(id, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *photoHandler) UpdatePhoto(ctx *gin.Context) {
	updatePhotoRequest := dto.UpdatePhotoRequest{}
	err := ctx.ShouldBind(&updatePhotoRequest)

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
	updatePhotoRequest.ID = id
	photo, err := c.photoService.UpdatePhoto(updatePhotoRequest, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", photo)
	ctx.JSON(http.StatusOK, response)

}
