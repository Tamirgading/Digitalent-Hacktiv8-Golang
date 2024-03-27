package controller

import (
	"MyGram/model"
	"MyGram/repository"
	"MyGram/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type photoController struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoController(photoRepository repository.PhotoRepository) *photoController {
	return &photoController{
		photoRepository: photoRepository,
	}
}

func (pc *photoController) CreatePhoto(ctx *gin.Context) {
	var newPhoto model.Photo
	if err := ctx.ShouldBindJSON(&newPhoto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := util.GetUserIDFromToken(ctx)

	newPhoto.UserID = userID

	if err := pc.photoRepository.Create(&newPhoto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newPhoto)
}

func (pc *photoController) GetAllPhotos(ctx *gin.Context) {
	photos, err := pc.photoRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (pc *photoController) UpdatePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoId")
	var updatePhoto model.Photo
	if err := ctx.ShouldBindJSON(&updatePhoto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photo, err := pc.photoRepository.GetById(photoID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	userID := util.GetUserIDFromToken(ctx)

	if photo.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this photo"})
		return
	}

	photo.Title = updatePhoto.Title
	photo.Caption = updatePhoto.Caption
	photo.PhotoURL = updatePhoto.PhotoURL

	if err := pc.photoRepository.Update(photo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func (pc *photoController) DeletePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoId")
	photo, err := pc.photoRepository.GetById(photoID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	userID := util.GetUserIDFromToken(ctx)

	if photo.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to delete this photo"})
		return
	}

	if err := pc.photoRepository.Delete(photoID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
