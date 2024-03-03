package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

package controllers

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarDates = []Car{}

func CreateCar(ctx *gin.Context){
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newCar.CarID = fmt.Sprintf("c%d", len(CarDates)+1)
	CarDates = append(CarDates, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}
func UpdateCar(ctx *gin.Context){
	carID := ctx.Param("carID")
	condition := false
	var updateCar Car

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDates {
		if carID == car.CarID{
			condition = true
			CarDates[i] = updateCar
			CarDates[i].CarID = carID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully updated", carID),
	})
}
