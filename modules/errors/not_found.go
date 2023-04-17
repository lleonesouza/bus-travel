package errors

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	Title   string `json:"title" example:"NotFoundError"`
	Details string `json:"details" example:"The email 'john@doe.com' is not found in the system."`
	Status  int    `json:"status" example:"404"`
}

func NotFound(entity string, value string) *NotFoundError {
	NotFound := &NotFoundError{
		Title:   "NotFoundError",
		Status:  http.StatusNotFound,
		Details: fmt.Sprintf("The %s '%s' is not found in the system.", entity, value),
	}

	return NotFound
}
