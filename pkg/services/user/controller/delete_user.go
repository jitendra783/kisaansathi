package controller

import (
	"context"
	"errors"
	"strings"

	"kisaanSathi/pkg/services/user/models"
)

func (c *controller) DeleteUser(
	ctx context.Context,
	email string,
) (*models.DeleteUserResponse, error) {

	email = strings.TrimSpace(strings.ToLower(email))

	if email == "" {
		return nil, errors.New("email is required")
	}

	// Check User Exists
	user, err := c.store.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// Soft Delete
	if err := c.store.DeleteUser(ctx, email); err != nil {
		return nil, err
	}

	return &models.DeleteUserResponse{
		Status:  true,
		Message: "User deleted successfully",
	}, nil
}