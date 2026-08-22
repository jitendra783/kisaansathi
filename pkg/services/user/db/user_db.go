package db

import (
	"context"
)

const (

	deleteUserQuery = `
		UPDATE users
		SET
			is_active = FALSE,
			updated_at = NOW()
		WHERE email = $1
	`

	activateUserQuery = `
		UPDATE users
		SET
			is_active = TRUE,
			updated_at = NOW()
		WHERE email = $1
	`

	deactivateUserQuery = `
		UPDATE users
		SET
			is_active = FALSE,
			updated_at = NOW()
		WHERE email = $1
	`
)

// =======================================================
// Delete User (Soft Delete)
// =======================================================

func (d *dbObject) DeleteUser(
	ctx context.Context,
	email string,
) error {

	_, err := d.db.ExecContext(
		ctx,
		deleteUserQuery,
		email,
	)

	return err
}

// =======================================================
// Activate User
// =======================================================

func (d *dbObject) ActivateUser(
	ctx context.Context,
	email string,
) error {

	_, err := d.db.ExecContext(
		ctx,
		activateUserQuery,
		email,
	)

	return err
}

// =======================================================
// Deactivate User
// =======================================================

func (d *dbObject) DeactivateUser(
	ctx context.Context,
	email string,
) error {

	_, err := d.db.ExecContext(
		ctx,
		deactivateUserQuery,
		email,
	)

	return err
}