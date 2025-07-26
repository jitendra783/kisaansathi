package db

import (
	"context"
	"errors"
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/common/models"
	"kisaanSathi/pkg/services/common/utils"
)

func (s schemeDB) commonSchemeValidation(ctx context.Context, schemeFlags *models.SchemeFlags, isOfflineRequest, isSpecialInterval bool) error {
	if isOfflineRequest && schemeFlags.OfflineFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.NotOfflineEnabled)
		return errors.New(utils.SchemeErrors.NotOfflineEnabled)
	}

	if !isOfflineRequest && schemeFlags.OnlineFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.NotOnlineEnabled)
		return errors.New(utils.SchemeErrors.NotOnlineEnabled)
	}

	return nil
}

func (s schemeDB) ValidateSchemeFlagsForPurchase(ctx context.Context, compCode int, schemeCode string, isOfflineRequest bool) (*models.SchemeFlags, error) {
	schemeFlags, err := s.GetSchemeFlags(ctx, compCode, schemeCode)
	if err != nil {
		return nil, err
	}

	err = s.commonSchemeValidation(ctx, schemeFlags, isOfflineRequest, false)
	if err != nil {
		return nil, err
	}

	if schemeFlags.PurchaseFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.PurchaseNotEnabled)
		return nil, errors.New(utils.SchemeErrors.PurchaseNotEnabled)
	}

	if schemeFlags.PurchaseAllowedFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.NoFurtherTransaction)
		return nil, errors.New(utils.SchemeErrors.NoFurtherTransaction)
	}

	return schemeFlags, nil
}

func (s schemeDB) ValidateSchemeFlagsForSIP(ctx context.Context, compCode int, schemeCode string, isOfflineRequest bool) (*models.SchemeFlags, error) {
	schemeFlags, err := s.GetSchemeFlags(ctx, compCode, schemeCode)
	if err != nil {
		return nil, err
	}

	err = s.commonSchemeValidation(ctx, schemeFlags, isOfflineRequest, false)
	if err != nil {
		return nil, err
	}

	if schemeFlags.SIPFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.SipNotEnabled)
		return nil, errors.New(utils.SchemeErrors.SipNotEnabled)
	}

	if schemeFlags.PurchaseAllowedFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.NoFurtherTransaction)
		return nil, errors.New(utils.SchemeErrors.NoFurtherTransaction)
	}

	return schemeFlags, nil
}

func (s schemeDB) ValidateSchemeFlagsForRedeem(ctx context.Context, compCode int, schemeCode string, isSpecialInterval bool) (*models.SchemeFlags, error) {
	schemeFlags, err := s.GetSchemeFlags(ctx, compCode, schemeCode)
	if err != nil {
		return nil, err
	}

	if schemeFlags.CloseFlag == "Close Ended" && schemeFlags.RedeemFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.RedeemCloseEnded)
		return nil, errors.New(utils.SchemeErrors.RedeemCloseEnded)
	}

	if schemeFlags.RedeemFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.ActivityNotAllowed)
		return nil, errors.New(utils.SchemeErrors.ActivityNotAllowed)
	}

	if schemeFlags.RedeemAllowedFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.NoFurtherTransaction)
		return nil, errors.New(utils.SchemeErrors.NoFurtherTransaction)
	}

	if isSpecialInterval && schemeFlags.SpecialIntervalFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.NotSpecialInterval)
		return nil, errors.New(utils.SchemeErrors.NotSpecialInterval)
	}

	return schemeFlags, nil
}

func (s schemeDB) ValidateSchemeFlagsForSWP(ctx context.Context, compCode int, schemeCode string, isSpecialInterval bool) (*models.SchemeFlags, error) {
	schemeFlags, err := s.GetSchemeFlags(ctx, compCode, schemeCode)
	if err != nil {
		return nil, err
	}

	if schemeFlags.SWPFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.ActivityNotAllowed)
		return nil, errors.New(utils.SchemeErrors.ActivityNotAllowed)
	}

	if schemeFlags.RedeemAllowedFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.NoFurtherTransaction)
		return nil, errors.New(utils.SchemeErrors.NoFurtherTransaction)
	}

	if isSpecialInterval && schemeFlags.SpecialIntervalFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.NotSpecialInterval)
		return nil, errors.New(utils.SchemeErrors.NotSpecialInterval)
	}

	return schemeFlags, nil
}

func (s schemeDB) ValidateSchemeFlagsForSwitch(ctx context.Context, compCode int, schemeCode string, isSpecialInterval bool) (*models.SchemeFlags, error) {
	schemeFlags, err := s.GetSchemeFlags(ctx, compCode, schemeCode)
	if err != nil {
		return nil, err
	}

	if schemeFlags.CloseFlag == "Close Ended" && schemeFlags.SwitchFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.SwitchCloseEnded)
		return nil, errors.New(utils.SchemeErrors.SwitchCloseEnded)
	}

	if schemeFlags.SwitchFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.ActivityNotAllowed)
		return nil, errors.New(utils.SchemeErrors.ActivityNotAllowed)
	}

	if schemeFlags.SwitchAllowedFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.NoFurtherTransaction)
		return nil, errors.New(utils.SchemeErrors.NoFurtherTransaction)
	}

	if isSpecialInterval && schemeFlags.SpecialIntervalFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.NotSpecialInterval)
		return nil, errors.New(utils.SchemeErrors.NotSpecialInterval)
	}

	return schemeFlags, nil
}

func (s schemeDB) ValidateSchemeFlagsForSTP(ctx context.Context, compCode int, schemeCode string, isSpecialInterval, isBoosterSTP, isDematHolding bool) (*models.SchemeFlags, error) {
	schemeFlags, err := s.GetSchemeFlags(ctx, compCode, schemeCode)
	if err != nil {
		return nil, err
	}

	if schemeFlags.CloseFlag == "Close Ended" && ((isBoosterSTP && schemeFlags.BoosterSTPFlag.IsNo()) || (!isBoosterSTP && schemeFlags.STPOutFlag.IsNo())) {
		logger.Log(ctx).Error(utils.SchemeErrors.StpCloseEnded)
		return nil, errors.New(utils.SchemeErrors.StpCloseEnded)
	}

	if isBoosterSTP && schemeFlags.BoosterSTPFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.BoosterSTPNotEnabled)
		return nil, errors.New(utils.SchemeErrors.BoosterSTPNotEnabled)
	}

	if !isBoosterSTP && schemeFlags.STPOutFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.StpNotEnabled)
		return nil, errors.New(utils.SchemeErrors.StpNotEnabled)
	}

	if schemeFlags.SwitchAllowedFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.NoFurtherTransaction)
		return nil, errors.New(utils.SchemeErrors.NoFurtherTransaction)
	}

	if isSpecialInterval && schemeFlags.SpecialIntervalFlag.IsNo() {
		logger.Log(ctx).Error(utils.SchemeErrors.NotSpecialInterval)
		return nil, errors.New(utils.SchemeErrors.NotSpecialInterval)
	}

	if isDematHolding {
		registrar, err := s.GetCompanyRegistrar(ctx, compCode)
		if err != nil || registrar == "KARVY" {
			logger.Log(ctx).Error(utils.SchemeErrors.StpNotAllowed)
			return nil, errors.New(utils.SchemeErrors.StpNotAllowed)
		}
	}

	return schemeFlags, nil
}
