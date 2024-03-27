package controller

import (
	"MyGram/model"
	"MyGram/repository"
	"MyGram/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) *userController {
	return &userController{
		userRepository: userRepository,
	}
}

func (uc *userController) RegisterUser(ctx *gin.Context) {
	var newUser model.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.userRepository.Create(&newUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newUser)
}

func (uc *userController) LoginUser(ctx *gin.Context) {
	var loginReq model.LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userRepository.GetByEmail(loginReq.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !util.CheckPasswordHash(loginReq.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := util.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("userId")
	var updateUser model.User
	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userRepository.GetById(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Email = updateUser.Email
	user.Username = updateUser.Username

	if err := uc.userRepository.Update(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *userController) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("userId")
	if err := uc.userRepository.Delete(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
