package controller

import (
	"MyGram/model"
	"MyGram/repository"
	"MyGram/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type socialMediaController struct {
	socialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaController(socialMediaRepository repository.SocialMediaRepository) *socialMediaController {
	return &socialMediaController{
		socialMediaRepository: socialMediaRepository,
	}
}

func (smc *socialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var newSocialMedia model.SocialMedia
	if err := ctx.ShouldBindJSON(&newSocialMedia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := util.GetUserIDFromToken(ctx)

	newSocialMedia.UserID = userID

	if err := smc.socialMediaRepository.Create(&newSocialMedia); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newSocialMedia)
}

func (smc *socialMediaController) GetAllSocialMedias(ctx *gin.Context) {
	socialMedias, err := smc.socialMediaRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"social_medias": socialMedias})
}

func (smc *socialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaId")
	var updateSocialMedia model.SocialMedia
	if err := ctx.ShouldBindJSON(&updateSocialMedia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	socialMedia, err := smc.socialMediaRepository.GetById(socialMediaID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Social media not found"})
		return
	}

	userID := util.GetUserIDFromToken(ctx)

	if socialMedia.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this social media"})
		return
	}

	socialMedia.Name = updateSocialMedia.Name
	socialMedia.SocialMediaURL = updateSocialMedia.SocialMediaURL

	if err := smc.socialMediaRepository.Update(socialMedia); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func (smc *socialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaId")
	socialMedia, err := smc.socialMediaRepository.GetById(socialMediaID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Social media not found"})
		return
	}

	userID := util.GetUserIDFromToken(ctx)

	if socialMedia.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to delete this social media"})
		return
	}

	if err := smc.socialMediaRepository.Delete(socialMediaID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Social media deleted successfully"})
}
