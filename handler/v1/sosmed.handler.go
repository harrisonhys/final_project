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

type SocialmediaHandler interface {
	All(ctx *gin.Context)
	CreateSocialmedia(ctx *gin.Context)
	UpdateSocialmedia(ctx *gin.Context)
	DeleteSocialmedia(ctx *gin.Context)
	FindOneSocialmediaByID(ctx *gin.Context)
}

type socialmediaHandler struct {
	socialmediaService service.SocialMediaService
	jwtService         service.JWTService
}

func NewSocialmediaHandler(socialmediaService service.SocialMediaService, jwtService service.JWTService) SocialmediaHandler {
	return &socialmediaHandler{
		socialmediaService: socialmediaService,
		jwtService:         jwtService,
	}
}

func (c *socialmediaHandler) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	socialmedias, err := c.socialmediaService.All(userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", socialmedias)
	ctx.JSON(http.StatusOK, response)
}

func (c *socialmediaHandler) CreateSocialmedia(ctx *gin.Context) {
	var createSocialmediaReq dto.CreateSocialMediaRequest
	err := ctx.ShouldBind(&createSocialmediaReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.socialmediaService.CreateSocialMedia(createSocialmediaReq, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *socialmediaHandler) FindOneSocialmediaByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.socialmediaService.FindOneSocialMediaByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *socialmediaHandler) DeleteSocialmedia(ctx *gin.Context) {
	id := ctx.Param("id")

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	err := c.socialmediaService.DeleteSocialMedia(id, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *socialmediaHandler) UpdateSocialmedia(ctx *gin.Context) {
	updateSocialmediaRequest := dto.UpdateSocialMediaRequest{}
	err := ctx.ShouldBind(&updateSocialmediaRequest)

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
	updateSocialmediaRequest.ID = id
	socialmedia, err := c.socialmediaService.UpdateSocialMedia(updateSocialmediaRequest, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", socialmedia)
	ctx.JSON(http.StatusOK, response)

}
