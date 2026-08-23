package db

import (
	"context"

	"kisaanSathi/pkg/services/user/models"

	"github.com/jmoiron/sqlx"
)

type dbObject struct {
	db *sqlx.DB
}

type UserStore interface {

	// ==========================
	// Authentication
	// ==========================

	CreateUser(ctx context.Context, user *models.User) error

	GetUserByID(ctx context.Context, id int64) (*models.User, error)

	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	IsEmailExists(ctx context.Context, email string) (bool, error)

	IsMobileExists(ctx context.Context, mobile string) (bool, error)

	// ==========================
	// Profile
	// ==========================

	UpdateUserDetails(
		ctx context.Context,
		req *models.UpdateUserRequest,
	) error

	UpdateAvatar(
		ctx context.Context,
		email string,
		avatar string,
	) error

	// ==========================
	// Password
	// ==========================

	UpdatePassword(
		ctx context.Context,
		email string,
		password string,
	) error

	// ==========================
	// User
	// ==========================

	DeleteUser(
		ctx context.Context,
		email string,
	) error

	// ActivateUser activates a user account by setting is_active to true.
	ActivateUser(
		ctx context.Context,
		email string,
	) error
	//deactivateUser deactivates a user account by setting is_active to false.
	DeactivateUser(
		ctx context.Context,
		email string,
	) error
}

func NewDBObject(db *sqlx.DB) UserStore {
	return &dbObject{
		db: db,
	}
}
