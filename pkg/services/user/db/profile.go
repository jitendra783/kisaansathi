package db

import (
	"context"

	"kisaanSathi/pkg/services/user/models"
)

const (
	updateUserDetailsQuery = `
		UPDATE users
		SET
			name       = :name,
			bio        = :bio,
			phone      = :mobile,
			address    = :address,
			state      = :state,
			district   = :district,
			zipcode    = :zipcode,
			language   = :language,
			soil_type  = :soil_type,
			updated_at = NOW()
		WHERE email = :email
		  AND is_active = TRUE
	`

	updateAvatarQuery = `
		UPDATE users
		SET
			avatar     = $1,
			updated_at = NOW()
		WHERE email = $2
		  AND is_active = TRUE
	`
)

// =======================================================
// Update User Profile
// =======================================================

func (d *dbObject) UpdateUserDetails(
	ctx context.Context,
	req *models.UpdateUserRequest,
) error {

	_, err := d.db.NamedExecContext(
		ctx,
		updateUserDetailsQuery,
		req,
	)

	return err
}

// =======================================================
// Update Avatar
// =======================================================

func (d *dbObject) UpdateAvatar(
	ctx context.Context,
	email string,
	avatar string,
) error {

	_, err := d.db.ExecContext(
		ctx,
		updateAvatarQuery,
		avatar,
		email,
	)

	return err
}
