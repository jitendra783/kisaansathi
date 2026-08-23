package db

import (
	"context"
	"database/sql"
	"errors"

	"kisaanSathi/pkg/services/user/models"
)

const (
	insertUserQuery = `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			phone,
			password,
			role,
			language,
			soil_type,
			state,
			district,
			address,
			zipcode
		)
		VALUES (
			:first_name,
			:last_name,
			:email,
			:phone,
			:password,
			:role,
			:language,
			:soil_type,
			:state,
			:district,
			:address,
			:zipcode
		)
	`

	selectUserByID = `
		SELECT
			*
		FROM users
		WHERE id = $1
		  AND is_active = TRUE
	`

	selectUserByEmail = `
		SELECT
			*
		FROM users
		WHERE email = $1
		  AND is_active = TRUE
	`

	emailExistsQuery = `
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE email = $1
		)
	`

	mobileExistsQuery = `
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE phone = $1
		)
	`
)

// ----------------------------------------------------
// Register
// ----------------------------------------------------

func (d *dbObject) CreateUser(
	ctx context.Context,
	user *models.User,
) error {

	_, err := d.db.NamedExecContext(
		ctx,
		insertUserQuery,
		user,
	)

	return err
}

// ----------------------------------------------------
// Login
// ----------------------------------------------------

func (d *dbObject) GetUserByID(
	ctx context.Context,
	id int64,
) (*models.User, error) {

	var user models.User

	err := d.db.GetContext(
		ctx,
		&user,
		selectUserByID,
		id,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (d *dbObject) GetUserByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {

	var user models.User

	err := d.db.GetContext(
		ctx,
		&user,
		selectUserByEmail,
		email,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

// ----------------------------------------------------
// Validation
// ----------------------------------------------------

func (d *dbObject) IsEmailExists(
	ctx context.Context,
	email string,
) (bool, error) {

	var exists bool

	err := d.db.GetContext(
		ctx,
		&exists,
		emailExistsQuery,
		email,
	)

	return exists, err
}

func (d *dbObject) IsMobileExists(
	ctx context.Context,
	mobile string,
) (bool, error) {

	var exists bool

	err := d.db.GetContext(
		ctx,
		&exists,
		mobileExistsQuery,
		mobile,
	)

	return exists, err
}
