package handler

import (
	"strconv"

	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/services/community/models"

	"github.com/gin-gonic/gin"
)

// GetPosts godoc
func (h *communityHandler) GetPosts(ctx *gin.Context) {

	resp, err := h.controller.GetPosts()
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()))
		return
	}

	network.SuccessResponse(resp)
}

// CreatePost godoc
func (h *communityHandler) CreatePost(ctx *gin.Context) {

	var req models.CreatePostRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()))
		return
	}

	resp, err := h.controller.CreatePost(&req)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()))
		return
	}

	network.SuccessResponse(resp)
}

// UpdatePost godoc
func (h *communityHandler) UpdatePost(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription("invalid post id"))
		return
	}

	var req models.UpdatePostRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()))
		return
	}

	resp, err := h.controller.UpdatePost(id, &req)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()))
		return
	}

	network.SuccessResponse(resp)
}

// DeletePost godoc
func (h *communityHandler) DeletePost(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription("invalid post id"))
		return
	}

	resp, err := h.controller.DeletePost(id)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()))
		return
	}

	network.SuccessResponse(resp)
}

// CreateComment godoc
func (h *communityHandler) CreateComment(ctx *gin.Context) {

	var req models.CreateCommentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()))
		return
	}

	resp, err := h.controller.CreateComment(&req)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()))
		return
	}

	network.SuccessResponse(resp)
}

// DeleteComment godoc
func (h *communityHandler) DeleteComment(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription("invalid comment id"))
		return
	}

	resp, err := h.controller.DeleteComment(id)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()))
		return
	}

	network.SuccessResponse(resp)
}

// LikePost godoc
func (h *communityHandler) LikePost(ctx *gin.Context) {

	var req models.LikeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()))
		return
	}

	resp, err := h.controller.LikePost(&req)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()))
		return
	}

	network.SuccessResponse(resp)
}

// SharePost godoc
func (h *communityHandler) SharePost(ctx *gin.Context) {

	var req models.ShareRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()))
		return
	}

	resp, err := h.controller.SharePost(&req)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()))
		return
	}

	network.SuccessResponse(resp)
}
