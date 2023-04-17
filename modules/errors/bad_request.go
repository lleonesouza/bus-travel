package errors

import (
	"net/http"
)

type BadRequestError struct {
	Title   string `json:"title" example:"BadRequestError"`
	Details string `json:"details" example:"Key: 'BusTravelCreateInput.Origin' Error:Field validation for 'Origin' failed on the 'required' tag"`
	Status  int    `json:"status" example:"400"`
}

func BadRequest(details string) *BadRequestError {
	BadRequest := &BadRequestError{
		Title:   "BadRequestError",
		Status:  http.StatusBadRequest,
		Details: details,
	}

	return BadRequest
}
