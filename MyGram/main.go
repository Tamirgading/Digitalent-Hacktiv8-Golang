package main

import (
	"MyGram/controller"
	"MyGram/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.AuthMiddleware())

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", controller.RegisterUser)
		userRoutes.POST("/login", controller.LoginUser)
		userRoutes.PUT("/:userId", controller.UpdateUser)
		userRoutes.DELETE("/:userId", controller.DeleteUser)
	}

	photoRoutes := r.Group("/photos")
	{
		photoRoutes.POST("", controller.CreatePhoto)
		photoRoutes.GET("", controller.GetPhotos)
		photoRoutes.PUT("/:photoId", controller.UpdatePhoto)
		photoRoutes.DELETE("/:photoId", controller.DeletePhoto)
	}

	commentRoutes := r.Group("/comments")
	{
		commentRoutes.POST("", controller.CreateComment)
		commentRoutes.GET("", controller.GetComments)
		commentRoutes.PUT("/:commentId", controller.UpdateComment)
		commentRoutes.DELETE("/:commentId", controller.DeleteComment)
	}

	socialMediaRoutes := r.Group("/socialmedias")
	{
		socialMediaRoutes.POST("", controller.CreateSocialMedia)
		socialMediaRoutes.GET("", controller.GetSocialMedias)
		socialMediaRoutes.PUT("/:socialMediaId", controller.UpdateSocialMedia)
		socialMediaRoutes.DELETE("/:socialMediaId", controller.DeleteSocialMedia)
	}

	r.Run(":8080")
}
