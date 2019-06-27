package errors

import (
	"fmt"
)

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Type       string `json:"type"`
}

func (e *Error) Error() string {
	if len(e.Message) == 0 {
		return fmt.Sprintf("%d %s", e.StatusCode, e.Type)
	}
	return fmt.Sprintf("%d %s %s", e.StatusCode, e.Type, e.Message)
}
