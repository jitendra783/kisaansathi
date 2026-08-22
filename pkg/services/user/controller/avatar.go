package controller

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"kisaanSathi/pkg/services/user/models"
)

var allowedImageTypes = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
}

const maxAvatarSize = 5 * 1024 * 1024 // 5 MB

func (c *controller) UpdateAvatar(
	ctx context.Context,
	email string,
	file *multipart.FileHeader,
) (*models.UpdateAvatarResponse, error) {

	email = strings.TrimSpace(strings.ToLower(email))

	if email == "" {
		return nil, errors.New("email is required")
	}

	if file == nil {
		return nil, errors.New("avatar file is required")
	}

	// Validate size
	if file.Size > maxAvatarSize {
		return nil, errors.New("avatar size should not exceed 5 MB")
	}

	// Validate extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedImageTypes[ext] {
		return nil, fmt.Errorf("unsupported file type: %s", ext)
	}

	// Check user
	exists, err := c.store.IsEmailExists(ctx, email)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("user not found")
	}

	// Upload file
	avatarURL, err := c.storage.SaveAvatar(file)
	if err != nil {
		return nil, err
	}

	// Update database
	if err := c.store.UpdateAvatar(ctx, email, avatarURL); err != nil {
		return nil, err
	}

	return &models.UpdateAvatarResponse{
		Status:  true,
		Message: "Avatar updated successfully",
		Avatar:  avatarURL,
	}, nil
}
