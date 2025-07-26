package network

import (
	"errors"
	"kisaanSathi/pkg/utils/validations"

	"github.com/go-playground/validator/v10"
)

type HttpResponse struct {
	Status       string `json:"status"`
	Data         any    `json:"data,omitempty"`
	Error        *Error `json:"error,omitempty"`
	ErrorMessage string `json:"FML_ERROR_MSG,omitempty"`
}

func SuccessResponse(data any) HttpResponse {

	return HttpResponse{
		Status: "success",
		Data:   data,
	}
}

func FailureResponse(error Error) HttpResponse {
	return HttpResponse{
		Status:       "failure",
		Error:        &error,
		ErrorMessage: error.Description,
	}
}

func BadRequestResponse(err error, request interface{}) HttpResponse {
	var validationErrs validator.ValidationErrors
	var badError Error
	if errors.As(err, &validationErrs) {
		badError = ApiErrors.BadRequest.WithErrorDescription(validations.GetCustomErrorMessages(validationErrs, request))
	} else {
		badError = ApiErrors.BadRequest.WithErrorDescription(err.Error())
	}

	return HttpResponse{
		Status:       "failure",
		Error:        &badError,
		ErrorMessage: badError.Description,
	}
}
