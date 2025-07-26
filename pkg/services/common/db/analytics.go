package db

import (
	"context"
	"kisaanSathi/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type analyticsDB struct {
	oracle *gorm.DB
}

type AnalyticsStore interface {
	UpdateCampaignClickEvent(ctx context.Context, matchAccount string) error
}

func NewAnalyticsStore(oracle *gorm.DB) AnalyticsStore {
	return &analyticsDB{oracle: oracle}
}

func (a analyticsDB) UpdateCampaignClickEvent(ctx context.Context, matchAccount string) error {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	err := a.oracle.Transaction(func(tx *gorm.DB) error {
		//tx.Create(&models.CampaignClick{MatchAccount: matchAccount, ClickDate: ""})
		err := tx.Exec(`INSERT INTO MCM_MF_CMPGN_MTCH VALUES (?, SYSDATE)`, matchAccount).Error
		if err != nil {
			logger.Log(ctx).Error("Campaign event - ", zap.String("Error", err.Error()))
		}
		return err
	})

	if err != nil {
		return err
	}

	return nil
}
