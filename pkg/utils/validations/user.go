package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var ValidateMatchAccount validator.Func = func(fl validator.FieldLevel) bool {
	// need confirmation: match account should start with 1 or 5 or 8
	// re := regexp.MustCompile(`^[158]\d{9}$`)

	re := regexp.MustCompile(`^[1-9]\d{9}$`)
	return re.MatchString(fl.Field().String())
}

var ValidateUserId validator.Func = func(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9#]+$`)
	return re.MatchString(fl.Field().String())
}
