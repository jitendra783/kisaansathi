package controller

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"kisaanSathi/pkg/services/user/models"
)

// ChangePassword changes the user's password.
func (c *controller) ChangePassword(
	ctx context.Context,
	req *models.ChangePasswordRequest,
) (*models.ChangePasswordResponse, error) {

	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	user, err := c.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// Verify old password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.OldPassword),
	); err != nil {
		return nil, errors.New("old password is incorrect")
	}

	// Prevent same password
	if req.OldPassword == req.NewPassword {
		return nil, errors.New("new password must be different")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	if err := c.store.UpdatePassword(
		ctx,
		req.Email,
		string(hash),
	); err != nil {
		return nil, err
	}

	return &models.ChangePasswordResponse{
		Status:  true,
		Message: "Password changed successfully",
	}, nil
}

// ForgotPassword sends password reset OTP/link.
func (c *controller) ForgotPassword(
	ctx context.Context,
	req *models.ForgotPasswordRequest,
) (*models.ForgotPasswordResponse, error) {

	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	exists, err := c.store.IsEmailExists(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("user not found")
	}

	// TODO:
	// Generate OTP
	// Save OTP in Redis/DB
	// Send Email/SMS

	return &models.ForgotPasswordResponse{
		Status:  true,
		Message: "Password reset link sent successfully",
	}, nil
}

// ResetPassword resets password using OTP/token.
func (c *controller) ResetPassword(
	ctx context.Context,
	req *models.ResetPasswordRequest,
) (*models.ResetPasswordResponse, error) {

	// TODO:
	// Validate OTP/Token
	// Fetch email
	// Update password

	return &models.ResetPasswordResponse{
		Status:  true,
		Message: "Password reset successfully",
	}, nil
}
