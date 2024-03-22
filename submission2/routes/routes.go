package routes

import (
	"net/http"
	"submission2/models"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/orders", CreateOrder)
	router.GET("/orders", GetOrders)
	router.PUT("/orders/:orderId", UpdateOrder)
	router.DELETE("/orders/:orderId", DeleteOrder)
}

func CreateOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.BindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, order)
}

func GetOrders(ctx *gin.Context) {
	var orders []models.Order
	ctx.JSON(http.StatusOK, orders)
}

func UpdateOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.BindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func DeleteOrder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
