package controller

import (
	"context"
	"errors"
	"strings"

	"kisaanSathi/pkg/services/user/models"
)

// GetUserDetails returns user profile.
func (c *controller) GetUserDetails(
	ctx context.Context,
	email string,
) (*models.UserDetailsResponse, error) {

	email = strings.TrimSpace(strings.ToLower(email))

	user, err := c.store.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	response := &models.UserDetailsResponse{
		Status:  true,
		Message: "Success",
		User: models.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Mobile:    user.Mobile,
			Role:      user.Role,
			Language:  user.Language,
			SoilType:  user.SoilType,
			State:     user.State,
			District:  user.District,
			Address:   user.Address,
			ZipCode:   user.ZipCode,
			Avatar:    user.Avatar,
			Bio:       user.Bio,
			Lat:       user.Lat,
			Lng:       user.Lng,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	return response, nil
}

// UpdateUserDetails updates user profile.
func (c *controller) UpdateUserDetails(
	ctx context.Context,
	req *models.UpdateUserRequest,
) (*models.UpdateUserResponse, error) {

	req.Name = strings.TrimSpace(req.Name)
	req.Bio = strings.TrimSpace(req.Bio)

	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	exists, err := c.store.IsEmailExists(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("user not found")
	}

	if err := c.store.UpdateUserDetails(ctx, req); err != nil {
		return nil, err
	}

	return &models.UpdateUserResponse{
		Status:  true,
		Message: "Profile updated successfully",
	}, nil
}