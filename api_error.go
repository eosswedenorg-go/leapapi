package leapapi

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type APIErrorDetail struct {
	Message string `json:"message"`
	File    string `json:"file"`
	Line    int64  `json:"line_number"`
	Method  string `json:"method"`
}

type APIErrorInner struct {
	Code    int64            `json:"code"`
	Name    string           `json:"name"`
	What    string           `json:"what"`
	Details []APIErrorDetail `json:"details"`
}

type APIError struct {
	Code    int64         `json:"code"`
	Message string        `json:"message"`
	Err     APIErrorInner `json:"error"`
}

func (e *APIError) IsEmpty() bool {
	return cmp.Equal(*e, APIError{})
}

func (e APIError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}
