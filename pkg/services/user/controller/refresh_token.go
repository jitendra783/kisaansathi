package controller

import (
	"context"
	"errors"
	"strconv"
	"strings"

	auth "kisaanSathi/pkg/middlewares"
	"kisaanSathi/pkg/services/user/models"
)

func (c *controller) RefreshToken(
	ctx context.Context,
	req *models.RefreshTokenRequest,
) (*models.RefreshTokenResponse, error) {

	req.RefreshToken = strings.TrimSpace(req.RefreshToken)

	if req.RefreshToken == "" {
		return nil, errors.New("refresh token is required")
	}

	// Validate Refresh Token
	claims, err := auth.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}
	userid, err := strconv.ParseInt(claims.UserID, 10, 64)
	// Fetch User
	user, err := c.store.GetUserByID(ctx, userid)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if !user.IsActive {
		return nil, errors.New("user account is inactive")
	}

	// Generate New Access Token

	accessToken, err := auth.GenerateJWT(
		claims.UserID,
		user.Email,
		user.Role,
	)
	if err != nil {
		return nil, err
	}

	return &models.RefreshTokenResponse{
		Status:  true,
		Message: "Token refreshed successfully",
		Token:   accessToken,
	}, nil
}
