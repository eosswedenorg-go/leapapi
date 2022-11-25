package leapapi

import (
	"testing"
)

func TestHTTPError_Error(t *testing.T) {
	type fields struct {
		Code    int
		Message string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "with message", fields: fields{Code: 200, Message: "OK"}, want: "server returned HTTP 200 OK"},
		{name: "without message", fields: fields{Code: 404}, want: "server returned HTTP 404 Not Found"},
		{name: "Unkown code", fields: fields{Code: 999}, want: "server returned HTTP 999"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := HTTPError{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("HTTPError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
