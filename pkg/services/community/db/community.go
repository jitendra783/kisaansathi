package db

import (
	"kisaanSathi/pkg/services/community/models"
)

// GetPosts returns all community posts.
func (s *communityStore) GetPosts() ([]models.Post, error) {

	var posts []models.Post

	query := `
	SELECT
		p.id,
		p.user_id,
		COALESCE(u.name,'') AS user_name,
		p.title,
		p.content,
		p.image_url,
		p.created_at,
		p.updated_at,
		(SELECT COUNT(*) FROM community_likes l WHERE l.post_id = p.id) AS likes,
		(SELECT COUNT(*) FROM community_comments c WHERE c.post_id = p.id) AS comments,
		(SELECT COUNT(*) FROM community_shares s WHERE s.post_id = p.id) AS shares
	FROM community_posts p
	LEFT JOIN users u
		ON u.id = p.user_id
	ORDER BY p.created_at DESC`

	err := s.db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// CreatePost inserts a new post.
func (s *communityStore) CreatePost(req *models.CreatePostRequest) error {

	query := `
	INSERT INTO community_posts
		(user_id,title,content,image_url,created_at,updated_at)
	VALUES
		(?,?,?,?,NOW(),NOW())`

	_, err := s.db.Exec(
		query,
		req.UserID,
		req.Title,
		req.Content,
		req.ImageURL,
	)

	return err
}

// UpdatePost updates an existing post.
func (s *communityStore) UpdatePost(id int64, req *models.UpdatePostRequest) error {

	query := `
	UPDATE community_posts
	SET
		title=?,
		content=?,
		image_url=?,
		updated_at=NOW()
	WHERE id=?`

	_, err := s.db.Exec(
		query,
		req.Title,
		req.Content,
		req.ImageURL,
		id,
	)

	return err
}

// DeletePost deletes a post.
func (s *communityStore) DeletePost(id int64) error {

	_, err := s.db.Exec(
		`DELETE FROM community_posts WHERE id=?`,
		id,
	)

	return err
}

// CreateComment inserts a comment.
func (s *communityStore) CreateComment(req *models.CreateCommentRequest) error {

	query := `
	INSERT INTO community_comments
		(post_id,user_id,comment,created_at)
	VALUES
		(?,?,?,NOW())`

	_, err := s.db.Exec(
		query,
		req.PostID,
		req.UserID,
		req.Comment,
	)

	return err
}

// DeleteComment deletes a comment.
func (s *communityStore) DeleteComment(id int64) error {

	_, err := s.db.Exec(
		`DELETE FROM community_comments WHERE id=?`,
		id,
	)

	return err
}

// LikePost stores a like.
func (s *communityStore) LikePost(req *models.LikeRequest) error {

	query := `
	INSERT INTO community_likes
		(post_id,user_id,created_at)
	VALUES
		(?,?,NOW())`

	_, err := s.db.Exec(
		query,
		req.PostID,
		req.UserID,
	)

	return err
}

// SharePost stores a share.
func (s *communityStore) SharePost(req *models.ShareRequest) error {

	query := `
	INSERT INTO community_shares
		(post_id,user_id,created_at)
	VALUES
		(?,?,NOW())`

	_, err := s.db.Exec(
		query,
		req.PostID,
		req.UserID,
	)

	return err
}
