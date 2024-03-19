package routers

import (
	"github.com/gin-gonic/gin"
	"sesi6-tes-gin/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)
	return router
}
