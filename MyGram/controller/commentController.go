package controller

import (
	"MyGram/model"
	"MyGram/repository"
	"MyGram/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentRepository repository.CommentRepository
}

func NewCommentController(commentRepository repository.CommentRepository) *commentController {
	return &commentController{
		commentRepository: commentRepository,
	}
}

func (cc *commentController) CreateComment(ctx *gin.Context) {
	var newComment model.Comment
	if err := ctx.ShouldBindJSON(&newComment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assume token contains user information, you need to implement this logic
	userID := util.GetUserIDFromToken(ctx)

	newComment.UserID = userID

	if err := cc.commentRepository.Create(&newComment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newComment)
}

func (cc *commentController) GetAllComments(ctx *gin.Context) {
	comments, err := cc.commentRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (cc *commentController) UpdateComment(ctx *gin.Context) {
	commentID := ctx.Param("commentId")
	var updateComment model.Comment
	if err := ctx.ShouldBindJSON(&updateComment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := cc.commentRepository.GetById(commentID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	userID := util.GetUserIDFromToken(ctx)

	if comment.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this comment"})
		return
	}

	comment.Message = updateComment.Message

	if err := cc.commentRepository.Update(comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func (cc *commentController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("commentId")
	comment, err := cc.commentRepository.GetById(commentID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	userID := util.GetUserIDFromToken(ctx)

	if comment.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to delete this comment"})
		return
	}

	if err := cc.commentRepository.Delete(commentID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
