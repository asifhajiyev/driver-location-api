package model

import (
	err "driver-location-api/error"
	"github.com/go-playground/validator"
)

type RestResponse struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
	ErrorDetails interface{} `json:"errorDetails,omitempty"`
}

var validate = validator.New()

func ValidateRequest(driverLocationRequest interface{}) []*err.FieldValidationError {
	var errors []*err.FieldValidationError
	e := validate.Struct(driverLocationRequest)
	if e != nil {
		for _, er := range e.(validator.ValidationErrors) {
			var element err.FieldValidationError
			element.FailedField = er.StructField()
			element.Tag = er.Tag()
			errors = append(errors, &element)
		}
	}
	return errors
}

func BuildRestResponse(code int, message string, data interface{}, errorDetails interface{}) *RestResponse {
	return &RestResponse{
		Code:         code,
		Message:      message,
		Data:         data,
		ErrorDetails: errorDetails,
	}
}
