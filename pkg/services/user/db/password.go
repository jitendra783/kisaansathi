package db

import (
	"context"
)

const updatePasswordQuery = `
	UPDATE users
	SET
		password   = $1,
		updated_at = NOW()
	WHERE email = $2
	  AND is_active = TRUE
`

// =======================================================
// Update Password
// =======================================================

func (d *dbObject) UpdatePassword(
	ctx context.Context,
	email string,
	password string,
) error {

	_, err := d.db.ExecContext(
		ctx,
		updatePasswordQuery,
		password,
		email,
	)

	return err
}
