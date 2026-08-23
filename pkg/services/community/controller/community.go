package controller

import (
	"kisaanSathi/pkg/services/community/models"
)

// GetPosts returns all community posts.
func (c *communityController) GetPosts() (*models.PostListResponse, error) {

	posts, err := c.store.GetPosts()
	if err != nil {
		return nil, err
	}

	return &models.PostListResponse{
		Data: posts,
	}, nil
}

// CreatePost creates a new post.
func (c *communityController) CreatePost(req *models.CreatePostRequest) (*models.CommonResponse, error) {

	err := c.store.CreatePost(req)
	if err != nil {
		return nil, err
	}

	return &models.CommonResponse{
		Message: "Post created successfully",
	}, nil
}

// UpdatePost updates an existing post.
func (c *communityController) UpdatePost(id int64, req *models.UpdatePostRequest) (*models.CommonResponse, error) {

	err := c.store.UpdatePost(id, req)
	if err != nil {
		return nil, err
	}

	return &models.CommonResponse{
		Message: "Post updated successfully",
	}, nil
}

// DeletePost deletes a post.
func (c *communityController) DeletePost(id int64) (*models.CommonResponse, error) {

	err := c.store.DeletePost(id)
	if err != nil {
		return nil, err
	}

	return &models.CommonResponse{
		Message: "Post deleted successfully",
	}, nil
}

// CreateComment creates a comment on a post.
func (c *communityController) CreateComment(req *models.CreateCommentRequest) (*models.CommonResponse, error) {

	err := c.store.CreateComment(req)
	if err != nil {
		return nil, err
	}

	return &models.CommonResponse{
		Message: "Comment added successfully",
	}, nil
}

// DeleteComment deletes a comment.
func (c *communityController) DeleteComment(id int64) (*models.CommonResponse, error) {

	err := c.store.DeleteComment(id)
	if err != nil {
		return nil, err
	}

	return &models.CommonResponse{
		Message: "Comment deleted successfully",
	}, nil
}

// LikePost likes a post.
func (c *communityController) LikePost(req *models.LikeRequest) (*models.CommonResponse, error) {

	err := c.store.LikePost(req)
	if err != nil {
		return nil, err
	}

	return &models.CommonResponse{
		Message: "Post liked successfully",
	}, nil
}

// SharePost shares a post.
func (c *communityController) SharePost(req *models.ShareRequest) (*models.CommonResponse, error) {

	err := c.store.SharePost(req)
	if err != nil {
		return nil, err
	}

	return &models.CommonResponse{
		Message: "Post shared successfully",
	}, nil
}
