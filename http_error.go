package leapapi

import (
	"fmt"
	"net/http"
	"strings"
)

type HTTPError struct {
	Code    int
	Message string
}

func (e HTTPError) Error() string {
	msg := e.Message
	if len(msg) < 1 {
		msg = http.StatusText(e.Code)
	}

	msg = fmt.Sprintf("server returned HTTP %d %s", e.Code, msg)
	return strings.TrimSpace(msg)
}
