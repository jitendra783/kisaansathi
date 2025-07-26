package validations

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var ValidateDDMMYYYY_1 validator.Func = func(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	//Parse the date in DD/MM/YYYY format
	_, err := time.Parse("02/01/2006", dateStr)
	return err == nil

}

// ValidateDDMonYYYY validates date in "24-Feb-2020" format
var ValidateDDMonYYYY_2 validator.Func = func(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	// Parse with exact DD-Mon-YYYY format
	_, err := time.Parse("02-Jan-2006", dateStr)
	return err == nil
}

var ValidateDDMMYYYY_2 validator.Func = func(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	//Parse the date in DD-MM-YYYY format
	_, err := time.Parse("02-01-2006", dateStr)
	return err == nil

}

// ValidateDDMonYYYY validates date in "24-Feb-2020" format
var ValidateDDMonYYYY_1 validator.Func = func(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	// Parse with exact DD/Mon/YYYY format
	_, err := time.Parse("02/Jan/2006", dateStr)
	return err == nil
}

// ValidateYYYYMMDD_1 validates date in "2020/02/24" format
var ValidateYYYYMMDD_1 validator.Func = func(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	// Parse with exact DD/Mon/YYYY format
	_, err := time.Parse("2006/01/02", dateStr)
	return err == nil
}

// ValidateYYYYMMDD_2 validates date in "2020-02-24" format
var ValidateYYYYMMDD_2 validator.Func = func(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	// Parse with exact DD/Mon/YYYY format
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}

// ValidateYYYYMonDD_1 validates date in "2006/Jan/02" format
var ValidateYYYYMonDD_1 validator.Func = func(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	// Parse with exact YYYY/Mon/DD format
	_, err := time.Parse("2006/Jan/02", dateStr)
	return err == nil
}

// ValidateYYYYMonDD_2 validates date in "2006-Jan-02" format
var ValidateYYYYMonDD_2 validator.Func = func(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	// Parse with exact YYYY-Mon-DD format
	_, err := time.Parse("2006-Jan-02", dateStr)
	return err == nil
}

// ValidateYYYY validates date in "2006" format
var ValidateYYYY validator.Func = func(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	_, err := time.Parse("2006", dateStr)
	return err == nil
}
