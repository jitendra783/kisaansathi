package utils

import (
	"kisaanSathi/pkg/utils/validations"

	"github.com/go-playground/validator/v10"
)

func RegisterValidations(v *validator.Validate) {
	// string validations
	v.RegisterValidation("alphanumws", validations.ValidateAlphanumericWithSpace)
	v.RegisterValidation("alphaws", validations.ValidateAlphabetWithSpace)
	v.RegisterValidation("alphanumOr_D", validations.ValidateAlphanumericOr_D)

	// user validations
	v.RegisterValidation("matchaccount", validations.ValidateMatchAccount)
	v.RegisterValidation("userid", validations.ValidateUserId)

	// scheme validations
	v.RegisterValidation("schemecode", validations.ValidateAlphanumericWithHyphen)
	v.RegisterValidation("isin", validations.ValidateIsin)

	//date validations

	//Parse the date in DD/MM/YYYY format
	v.RegisterValidation("ddmmyyyy_1", validations.ValidateDDMMYYYY_1)

	//Parse the date in DD-MM-YYYY format
	v.RegisterValidation("ddmmyyyy_2", validations.ValidateDDMMYYYY_2)

	// Parse with exact DD/Mon/YYYY format
	v.RegisterValidation("ddmonyyyy_1", validations.ValidateDDMonYYYY_1)

	// Parse with exact DD-Mon-YYYY format
	v.RegisterValidation("ddmonyyyy_2", validations.ValidateDDMonYYYY_2)

	// Parse with exact YYYY/MM/DD format
	v.RegisterValidation("yyyymmdd_1", validations.ValidateYYYYMMDD_1)

	// Parse with exact YYYY-MM-DD format
	v.RegisterValidation("yyyymmdd_2", validations.ValidateYYYYMMDD_2)

	// Parse with exact YYYY/Mon/DD format
	v.RegisterValidation("yyyymondd_1", validations.ValidateYYYYMonDD_1)

	// Parse with exact YYYY-Mon-DD format
	v.RegisterValidation("yyyymondd_2", validations.ValidateYYYYMonDD_2)

	// Parse with exact YYYY format
	v.RegisterValidation("yyyy", validations.ValidateYYYY)
}
