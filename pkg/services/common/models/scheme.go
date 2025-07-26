package models

import "kisaanSathi/pkg/services/common/utils"

type SchemeFlags struct {
	BoosterSTPFlag        utils.Strings
	CloseFlag             utils.Strings
	DirectSchemeFlag      utils.Strings
	DRFlag                utils.Strings
	DivReinvestFlag       utils.Strings
	ETFFlag               utils.Strings
	FreedomFlag           utils.Strings
	FreeInsureFlag        utils.Strings
	MultiTransAllowedFlag utils.Strings
	OfflineFlag           utils.Strings
	OnlineFlag            utils.Strings
	PurchaseAllowedFlag   utils.Strings
	PurchaseFlag          utils.Strings
	RecommendFlag         utils.Strings
	RedeemFlag            utils.Strings
	RedeemAllowedFlag     utils.Strings
	RenewalFlag           utils.Strings
	SIPFlag               utils.Strings
	SpecialIntervalFlag   utils.Strings
	StepUpFlag            utils.Strings
	STPOutFlag            utils.Strings
	SwitchFlag            utils.Strings
	SwitchAllowedFlag     utils.Strings
	SWPFlag               utils.Strings
	TargetFundFlag        utils.Strings
}

type NavDetails struct {
	AmfiCode string
	NavDate  string
	NavValue float64
}

type SchemeDetails struct {
	MinPurchaseAmount   float64
	MinSIPAmount        float64
	MultiPurchaseAmount float64
	MultiSIPAmount      float64
	MaxSipAmount        float64
	MaxSubAmount        float64
	NFOExecDate         string
	SchemeDesc          string
	SchemeType          string
}

type RedeemSchemeDetails struct {
	AmfiCode               string
	NavDate                string
	NavValue               float64
	RedeemCutoffTime       string
	SisoCutoffTime         string
	MinRedeemAmount        float64
	MultiRedeemAmount      float64
	MinRedeemUnits         float64
	MultiRedeemUnits       float64
	StpMinAmount           float64
	StpMaxAmount           float64
	StpMinHoldingAmount    float64
	AccountTypeAllowedFlag string
}
