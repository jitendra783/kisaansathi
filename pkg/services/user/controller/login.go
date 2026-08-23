package controller

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"

	auth "kisaanSathi/pkg/middlewares"
	"kisaanSathi/pkg/services/user/models"
)

func (c *controller) Login(
	ctx context.Context,
	req *models.LoginRequest,
) (*models.LoginResponse, error) {

	// Normalize input
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	// Fetch user
	user, err := c.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check active user
	if !user.IsActive {
		return nil, errors.New("user account is inactive")
	}

	// Generate JWT
	token, err := auth.GenerateJWT(
		strconv.Itoa(user.ID),
		user.Email,
		user.Role,
	)
	if err != nil {
		return nil, err
	}

	response := &models.LoginResponse{
		Token:        token,
		RefreshToken: token,
	}

	return response, nil
}
