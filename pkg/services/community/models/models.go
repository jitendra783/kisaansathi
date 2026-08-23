package models

import "time"

// ==============================
// Request Models
// ==============================

// CreatePostRequest represents the request to create a new community post.
type CreatePostRequest struct {
	UserID   int64  `json:"user_id" binding:"required"`
	Title    string `json:"title"`
	Content  string `json:"content" binding:"required"`
	ImageURL string `json:"image_url"`
}

// UpdatePostRequest represents the request to update a community post.
type UpdatePostRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content" binding:"required"`
	ImageURL string `json:"image_url"`
}

// CreateCommentRequest represents the request to add a comment.
type CreateCommentRequest struct {
	PostID  int64  `json:"post_id" binding:"required"`
	UserID  int64  `json:"user_id" binding:"required"`
	Comment string `json:"comment" binding:"required"`
}

// LikeRequest represents the request to like a post.
type LikeRequest struct {
	PostID int64 `json:"post_id" binding:"required"`
	UserID int64 `json:"user_id" binding:"required"`
}

// ShareRequest represents the request to share a post.
type ShareRequest struct {
	PostID int64 `json:"post_id" binding:"required"`
	UserID int64 `json:"user_id" binding:"required"`
}

// ==============================
// Response Models
// ==============================

// Post represents a community post.
type Post struct {
	ID        int64     `db:"id" json:"id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	UserName  string    `db:"user_name" json:"user_name"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	ImageURL  string    `db:"image_url" json:"image_url"`
	Likes     int       `json:"likes"`
	Comments  int       `json:"comments"`
	Shares    int       `json:"shares"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Comment represents a post comment.
type Comment struct {
	ID        int64     `db:"id" json:"id"`
	PostID    int64     `db:"post_id" json:"post_id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	UserName  string    `db:"user_name" json:"user_name"`
	Comment   string    `db:"comment" json:"comment"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// GetPostsResponse represents the response for listing posts.
type GetPostsResponse struct {
	Data []Post `json:"data"`
}

// CommonResponse represents a generic success response.
type CommonResponse struct {
	Message string `json:"message"`
}

type PostListResponse struct {
	Data       []Post `json:"data"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}
