package db

import (
	"context"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/banner/models"
)

func (d *dbObject) GetBanners(ctx context.Context) ([]*models.Banner, error) {

	logger.Log().Info("Banner DB: GetBanners started")

	query := `
		SELECT
			id,
			page,
			title,
			subtitle,
			image_url,
			action_url,
			display_order,
			is_active
		FROM kisansathi.ui_banners
		WHERE is_active = TRUE
		ORDER BY display_order ASC, id ASC;
	`

	banners := make([]*models.Banner, 0)

	err := d.db.SelectContext(ctx, &banners, query)
	if err != nil {
		logger.Log().Error("Banner DB: " + err.Error())
		return nil, err
	}
	for i, banner := range banners {
		logger.Log().Sugar().Infof("Banner %d: %+v", i+1, *banner)
	}
	logger.Log().Sugar().Infof("Banner DB: Total banners fetched: %d", len(banners))

	return banners, nil
}

func (d *dbObject) GetBannerByID(ctx context.Context, id int64) (*models.Banner, error) {

	logger.Log().Sugar().Infof("Banner DB: GetBannerByID called. ID=%d", id)

	query := `
		SELECT
			id,
			page,
			title,
			subtitle,
			image_url,
			action_url,
			display_order,
			is_active
		FROM kisansathi.ui_banners
		WHERE id = $1;
	`

	var banner models.Banner

	err := d.db.GetContext(ctx, &banner, query, id)
	if err != nil {
		logger.Log().Error("Banner DB: " + err.Error())
		return nil, err
	}

	logger.Log().Info("Banner DB: Banner fetched successfully")

	return &banner, nil
}

func (d *dbObject) GetActiveBanners(ctx context.Context) ([]*models.Banner, error) {

	logger.Log().Info("Banner DB: GetActiveBanners started")

	query := `
		SELECT
			id,
			page,
			title,
			subtitle,
			image_url,
			action_url,
			display_order,
			is_active
		FROM kisansathi.ui_banners
		WHERE is_active = TRUE
		ORDER BY display_order ASC, id ASC;
	`

	banners := make([]*models.Banner, 0)

	err := d.db.SelectContext(ctx, &banners, query)
	if err != nil {
		logger.Log().Error("Banner DB: " + err.Error())
		return nil, err
	}

	logger.Log().Sugar().Infof("Banner DB: Active banners fetched: %d", len(banners))

	return banners, nil
}

func (d *dbObject) CreateBanner(ctx context.Context, banner *models.Banner) error {

	logger.Log().Info("Banner DB: CreateBanner started")

	query := `
		INSERT INTO kisansathi.ui_banners (
			page,
			title,
			subtitle,
			image_url,
			action_url,
			display_order,
			is_active
		)
		VALUES (
			:page,
			:title,
			:subtitle,
			:image_url,
			:action_url,
			:display_order,
			:is_active
		);
	`

	_, err := d.db.NamedExecContext(ctx, query, banner)
	if err != nil {
		logger.Log().Error("Banner DB: " + err.Error())
		return err
	}

	logger.Log().Info("Banner DB: Banner created successfully")

	return nil
}

func (d *dbObject) UpdateBanner(ctx context.Context, banner *models.Banner) error {

	logger.Log().Sugar().Infof("Banner DB: UpdateBanner called. ID=%d", banner.ID)

	query := `
		UPDATE kisansathi.ui_banners
		SET
			page = :page,
			title = :title,
			subtitle = :subtitle,
			image_url = :image_url,
			action_url = :action_url,
			display_order = :display_order,
			is_active = :is_active
		WHERE id = :id;
	`

	_, err := d.db.NamedExecContext(ctx, query, banner)
	if err != nil {
		logger.Log().Error("Banner DB: " + err.Error())
		return err
	}

	logger.Log().Info("Banner DB: Banner updated successfully")

	return nil
}

func (d *dbObject) DeleteBanner(ctx context.Context, id int64) error {

	logger.Log().Sugar().Infof("Banner DB: DeleteBanner called. ID=%d", id)

	query := `
		DELETE
		FROM kisansathi.ui_banners
		WHERE id = $1;
	`

	_, err := d.db.ExecContext(ctx, query, id)
	if err != nil {
		logger.Log().Error("Banner DB: " + err.Error())
		return err
	}

	logger.Log().Info("Banner DB: Banner deleted successfully")

	return nil
}
