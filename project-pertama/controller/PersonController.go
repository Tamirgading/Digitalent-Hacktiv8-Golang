package controller

import (
	"fmt"
	"net/http"
	_ "project-pertama/lib"
	"project-pertama/model"
	"project-pertama/repository"
	"project-pertama/util"
	_ "strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/google/uuid"
	_ "github.com/pelletier/go-toml/query"
)

type personController struct {
	personRepository repository.IPersonRepository
}

func NewPersonController(personRepository repository.IPersonRepository) *personController{
	return &personController{
		personRepository: personRepository,
	}
}

func (pc *personController) Create(ctx *gin.Context){
	var newPerson model.Person

	err := ctx.ShouldBindJSON(&newPerson)
	if err != nil{
		var r model.Response = model.Response{
			Success: false,
			Error: err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	createdPerson, err := pc.personRepository.Create(newPerson)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, createdPerson, ""))
}

// GetAll Person godoc
// @Summary      Get All Person
// @Schemes
// @Description  get all person
// @Tags         person
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.Person
// @Router       /person [get]
func (pc *personController) GetAll(ctx *gin.Context){
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.JSON(http.StatusUnauthorized, util.CreateResponse(false, nil, "unauthorized"))
		return
	}
	if username != "admin" || password != "12345"{
		ctx.JSON(http.StatusUnauthorized, util.CreateResponse(false, nil, "Wrong username / password"))
		return
	}
	persons, err := pc.personRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.CreateResponse(true, persons, ""))
}
func (pc *personController) Delete(ctx *gin.Context){
	idString := ctx.Param("id")
	fmt.Println(idString)
	err := pc.personRepository.Delete(idString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.CreateResponse(true, nil, ""))
}