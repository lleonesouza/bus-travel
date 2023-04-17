package errors

import (
	"net/http"
)

type InternalServerError struct {
	Title   string `json:"title" example:"InternalServerError"`
	Details string `json:"details" example:"Something goes wrong. Please try again"`
	Status  int    `json:"status" example:"500"`
}

func InternalServer(details string) *InternalServerError {
	InternalServer := &InternalServerError{
		Title:   "InternalServerError",
		Status:  http.StatusInternalServerError,
		Details: details,
	}

	return InternalServer
}
