package utils

type schemeError struct {
	NoSchemeFound             string
	NoRegistrarFound          string
	NoCompanyFound            string
	NoNavFound                string
	InvalidTransactionCode    string
	NoFurtherTransaction      string
	NotOfflineEnabled         string
	NotOnlineEnabled          string
	BoosterSTPNotEnabled      string
	NotSpecialInterval        string
	PurchaseNotEnabled        string
	SipNotEnabled             string
	StpNotAllowed             string
	StpNotEnabled             string
	ActivityNotAllowed        string
	InvalidNFOExecDate        string
	InvalidSIPPeriod          string
	SIPDatesNotFound          string
	SIPDateError              string
	SIPFrequencyNotFound      string
	SIPMinMultiAmountNotFound string
	RedeemCloseEnded          string
	SwitchCloseEnded          string
	StpCloseEnded             string
	FutureDateNotAvailable    string
	SWPDatesNotFound          string
	NoUserFound               string
}

var SchemeErrors = newSchemeErrorsRegistry()

func newSchemeErrorsRegistry() *schemeError {
	return &schemeError{
		NoSchemeFound:             "No scheme found",
		NoRegistrarFound:          "No registrar found for the company",
		NoCompanyFound:            "No company found",
		NoNavFound:                "No nav found",
		InvalidTransactionCode:    "Invalid transaction code",
		NoFurtherTransaction:      "The scheme is closed for further transaction",
		NotOfflineEnabled:         "Scheme not offline enabled",
		NotOnlineEnabled:          "Scheme not online enabled",
		BoosterSTPNotEnabled:      "This Scheme is not enabled for time the market STP",
		NotSpecialInterval:        "Scheme is not Special Interval",
		PurchaseNotEnabled:        "This scheme is not enabled for Purchase",
		SipNotEnabled:             "This scheme is not enabled for SIP",
		StpNotAllowed:             "STP order not allowed",
		StpNotEnabled:             "Scheme not enabled for STP Out",
		ActivityNotAllowed:        "Requested activity not allowed on this scheme",
		InvalidNFOExecDate:        "NFO Exec date is null",
		InvalidSIPPeriod:          "SIP Period is null",
		SIPDatesNotFound:          "SIP Dates not found",
		SWPDatesNotFound:          "SWP Dates not found",
		SIPDateError:              "Start date is less than or equal to today's date",
		SIPFrequencyNotFound:      "SIP Frequency not found",
		SIPMinMultiAmountNotFound: "SIP Min/Multi Amount not found",
		RedeemCloseEnded:          "This scheme is close ended scheme and currently not available for redemption",
		SwitchCloseEnded:          "This scheme is close ended scheme and currently not available for switch out",
		StpCloseEnded:             "This scheme is close ended scheme and currently not available for STP out",
		FutureDateNotAvailable:    "Future Date is not available for applied scheme",
		NoUserFound:               "No user found",
	}
}
