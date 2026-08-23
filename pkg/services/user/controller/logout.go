package controller

import (
	"context"
	"errors"
	"strings"

	"kisaanSathi/pkg/services/user/models"
)

func (c *controller) Logout(
	ctx context.Context,
	req *models.LogoutRequest,
) (*models.LogoutResponse, error) {

	req.LogoutFlag = strings.TrimSpace(strings.ToUpper(req.LogoutFlag))

	if req.LogoutFlag != "Y" && req.LogoutFlag != "N" {
		return nil, errors.New("invalid logout flag")
	}

	// ---------------------------------------------------
	// Future Enhancement:
	// 1. Extract JWT from Authorization header.
	// 2. Store token in Redis blacklist.
	// 3. Set expiry equal to JWT expiry.
	// ---------------------------------------------------

	response := &models.LogoutResponse{
		Status:  true,
		Message: "Logout successful",
	}

	return response, nil
}
