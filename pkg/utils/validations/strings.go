package validations

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Included by default
// alpha, alphanum, number, numeric, boolean, startswith, endswith, uppercase, lowercase

var ValidateAlphanumericWithSpace validator.Func = func(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	return re.MatchString(fl.Field().String())
}

var ValidateAlphanumericWithHyphen validator.Func = func(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	return re.MatchString(fl.Field().String())
}

var ValidateAlphabetWithSpace validator.Func = func(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	return re.MatchString(fl.Field().String())
}

var ValidateAlphanumericOr_D validator.Func = func(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if val == "_D" {
		return true
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return re.MatchString(val)
}

var ValidateIsin validator.Func = func(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^IN[EF][0-9A-Z]{8,9}[0-9]$`)
	return re.MatchString(fl.Field().String())
}

// Custom function to extract error messages from struct tags
func GetCustomErrorMessages(err error, obj interface{}) string {
	errorMessage := ""
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		objType := reflect.TypeOf(obj)
		fieldErr := validationErrors[0]
		field, _ := objType.FieldByName(fieldErr.StructField())
		jsonTag := field.Tag.Get("json")
		errorTag := field.Tag.Get("error")

		if jsonTag != "" {
			errorMessage = strings.Split(jsonTag, ",")[0] + " is invalid"
		} else {
			errorMessage = fieldErr.Field() + " is invalid"
		}

		if errorTag != "" {
			errorMessage = errorMessage + " - " + errorTag
		}
	}
	return errorMessage
}
