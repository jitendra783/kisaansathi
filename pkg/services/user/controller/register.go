package controller

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"kisaanSathi/pkg/services/user/models"
)

func (c *controller) Register(ctx context.Context, req *models.RegisterRequest) error {

	// Trim Spaces
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Mobile = strings.TrimSpace(req.Mobile)
	req.FirstName = strings.TrimSpace(req.FirstName)
	req.LastName = strings.TrimSpace(req.LastName)

	// Check Email
	exists, err := c.store.IsEmailExists(ctx, req.Email)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("email already registered")
	}

	// Check Mobile
	exists, err = c.store.IsMobileExists(ctx, req.Mobile)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("mobile number already registered")
	}

	// Encrypt Password
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Mobile:    req.Mobile,
		Password:  string(hash),
		Address:   req.Address,
		State:     req.State,
		District:  req.District,
		ZipCode:   req.ZipCode,
		Language:  req.Language,
		SoilType:  req.SoilType,
		Role:      "farmer",
	}

	return c.store.CreateUser(ctx, &user)
}