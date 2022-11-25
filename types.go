package leapapi

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/imroc/req/v3"
)

type Client struct {
	Url    string
	Host   string
	client *req.Client
}

// get_info format (not all fields).
type Info struct {
	ServerVersion  string    `json:"server_version"`
	HeadBlockNum   int64     `json:"head_block_num"`
	HeadBlockTime  time.Time `json:"head_block_time"`
	HTTPStatusCode int
}

// Service struct from /v2/health
type Service struct {
	Name   string
	Status string
	Data   map[string]interface{}
	Time   time.Time
}

func (s *Service) UnmarshalJSON(b []byte) error {
	var r struct {
		N string                 `json:"service"`
		S string                 `json:"status"`
		D map[string]interface{} `json:"service_data"`
		T int64                  `json:"time"`
	}

	err := json.Unmarshal(b, &r)
	if err == nil {
		s.Name = r.N
		s.Status = r.S
		s.Data = r.D
		s.Time = fromTS(r.T)
	}
	return err
}

// /v2/health format
type Health struct {
	Version        string                 `json:"version"`
	VersionHash    string                 `json:"version_hash"`
	Health         []Service              `json:"health"`
	Features       map[string]interface{} `json:"features"`
	QueryTime      float32                `json:"query_time_ms"`
	HTTPStatusCode int
}

// Error
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
	return HTTPError{Code: int(e.Code), Message: e.Message}.Error()
}
