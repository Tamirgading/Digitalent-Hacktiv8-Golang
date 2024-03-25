package main

import (
	"project-pertama/controller"
	"project-pertama/lib"
	"project-pertama/model"
	"project-pertama/repository"

	"github.com/gin-gonic/gin"
)

func main(){

	db, err := lib.InitDatabase()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Person{}, &model.CreditCard{})
	if err != nil {
		panic(err)
	}
	personRepository := repository.NewPersonRepository(db)
	personController := controller.NewPersonController(personRepository)

	ginEngine := gin.Default()

	ginEngine.GET("/person", personController.GetAll)
	ginEngine.POST("/person", personController.Create)
	ginEngine.DELETE("/person/:id", personController.Delete)

	err = ginEngine.Run("localhost:8082")
	if err != nil {
		panic(err)
	}
}