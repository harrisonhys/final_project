package main

import (
	"final_project/config"
	v1 "final_project/handler/v1"
	"final_project/middleware"
	"final_project/repo"
	"final_project/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.SetupDatabaseConnection()
	userRepo       repo.UserRepository        = repo.NewUserRepo(db)
	photoRepo      repo.PhotoRepository       = repo.NewPhotoRepo(db)
	commentRepo    repo.CommentRepository     = repo.NewCommentRepo(db)
	sosmedRepo     repo.SocialmediaRepository = repo.NewSocialmediaRepo(db)
	authService    service.AuthService        = service.NewAuthService(userRepo)
	jwtService     service.JWTService         = service.NewJWTService()
	userService    service.UserService        = service.NewUserService(userRepo)
	photoService   service.PhotoService       = service.NewPhotoService(photoRepo)
	commentService service.CommentService     = service.NewCommentService(commentRepo)
	sosmedService  service.SocialMediaService = service.NewSocialMediaService(sosmedRepo)
	authHandler    v1.AuthHandler             = v1.NewAuthHandler(authService, jwtService, userService)
	userHandler    v1.UserHandler             = v1.NewUserHandler(userService, jwtService)
	photoHandler   v1.PhotoHandler            = v1.NewPhotoHandler(photoService, jwtService)
	commentHandler v1.CommentHandler          = v1.NewCommentHandler(commentService, jwtService)
	sosmedHandler  v1.SocialmediaHandler      = v1.NewSocialmediaHandler(sosmedService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	photoRoutes := server.Group("api/photo", middleware.AuthorizeJWT(jwtService))
	{
		photoRoutes.GET("/", photoHandler.All)
		photoRoutes.POST("/", photoHandler.CreatePhoto)
		photoRoutes.GET("/:id", photoHandler.FindOnePhotoByID)
		photoRoutes.PUT("/:id", photoHandler.UpdatePhoto)
		photoRoutes.DELETE("/:id", photoHandler.DeletePhoto)
	}

	commentRoutes := server.Group("api/comment", middleware.AuthorizeJWT(jwtService))
	{
		commentRoutes.GET("/", commentHandler.All)
		commentRoutes.POST("/", commentHandler.CreateComment)
		commentRoutes.GET("/:id", commentHandler.FindOneCommentByID)
		commentRoutes.PUT("/:id", commentHandler.UpdateComment)
		commentRoutes.DELETE("/:id", commentHandler.DeleteComment)
	}

	sosmedRoutes := server.Group("api/sosmed", middleware.AuthorizeJWT(jwtService))
	{
		sosmedRoutes.GET("/", sosmedHandler.All)
		sosmedRoutes.POST("/", sosmedHandler.CreateSocialmedia)
		sosmedRoutes.GET("/:id", sosmedHandler.FindOneSocialmediaByID)
		sosmedRoutes.PUT("/:id", sosmedHandler.UpdateSocialmedia)
		sosmedRoutes.DELETE("/:id", sosmedHandler.DeleteSocialmedia)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
	}

	server.Run()
}
