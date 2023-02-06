package leapapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
	if req.URL.String() == "/v1/chain/get_info" {
		info := `{
            "server_version": "d1bc8d3",
            "head_block_num": 8888,
            "head_block_time": "2018-01-01T13:37:01"
        }`

		_, _ = res.Write([]byte(info))

	}

	if req.URL.String() == "/v2/health" {
		info := `{
            "version": "1.0",
            "version_hash": "028d5a34463884fcbe2ecfd3c0fcb3b5d4d538f4fd64803c1ef7209c85f2f266",
            "host": "api.test.com:443",
            "health": [
                {
                    "service": "Service1",
                    "status": "OK",
                    "time": 1642174781678
                },
                {
                    "service": "Service2",
                    "status": "DOWN",
                    "service_data": {
                        "key1": 1234,
                        "key2": "some_string"
                    },
                    "time": 1642174781736
                }
            ]
        }`

		_, _ = res.Write([]byte(info))
	}
}))

func TestSendContextTimeout(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Second * 4)
	}))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	client := New(testServer.URL)

	err := client.send(ctx, "GET", "/", nil, nil)
	assert.Error(t, err)
	assert.True(t, strings.HasSuffix(err.Error(), "deadline exceeded"), "Error was not deadline exceeded")
}

func TestSendContextCancel(t *testing.T) {
	done := make(chan interface{})
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Second * 10)
	}))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	client := New(testServer.URL)

	go func() {
		defer close(done)
		err := client.send(ctx, "GET", "/", nil, nil)
		assert.Error(t, err)
		assert.True(t, strings.HasSuffix(err.Error(), "context canceled"), "Error was not context canceled")
	}()

	time.Sleep(time.Second)
	cancel()

	<-done
}

func TestSendUrlParseFails(t *testing.T) {
	client := New("api.mylittleponies.org\n")

	err := client.send(context.Background(), "GET", "/v1/ponies/Rainbow Dash", nil, nil)
	assert.EqualError(t, err, "parse \"api.mylittleponies.org\\n\": net/url: invalid control character in URL")
}

func TestSendDefaultHostHeader(t *testing.T) {
	expected := ""

	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, expected, req.Host)
	}))

	u, err := url.Parse(srv.URL)
	require.NoError(t, err)

	expected = u.Host
	client := New(srv.URL)

	err = client.send(context.Background(), "GET", "/", nil, nil)
	require.NoError(t, err)
}

func TestSendCustomHostHeader(t *testing.T) {
	expected := "CustomHost"

	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, expected, req.Host)
	}))

	client := New(srv.URL)
	client.Host = expected

	err := client.send(context.Background(), "GET", "/", nil, nil)
	require.NoError(t, err)
}

func TestGetInfo(t *testing.T) {
	client := New(testServer.URL)

	info, err := client.GetInfo(context.Background())

	require.NoError(t, err)
	assert.Equal(t, "d1bc8d3", info.ServerVersion)
	assert.Equal(t, int64(8888), info.HeadBlockNum)
	assert.Equal(t, time.Unix(1514813821, 0).UTC(), info.HeadBlockTime)
}

func TestGetInfoHTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		payload := `{
            "code":500,
            "message":"Internal Server Error",
            "error":{
                "code":22,
                "name":"assertion",
                "what":"unspecified",
                "details":[
                    {
                        "message":"Assertion failed: a != b",
                        "file":"abi_reader.cpp",
                        "line_number":271,
                        "method":"read_abi"
                    }
                ]
            }
        }`
		res.WriteHeader(500)
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	_, err := client.GetInfo(context.Background())
	require.EqualError(t, err, "500 Internal Server Error")

	api_err, ok := err.(APIError)
	require.True(t, ok)

	expected := APIError{
		Code:    500,
		Message: "Internal Server Error",
		Err: APIErrorInner{
			Code: 22,
			Name: "assertion",
			What: "unspecified",
			Details: []APIErrorDetail{
				{
					Message: "Assertion failed: a != b",
					File:    "abi_reader.cpp",
					Line:    271,
					Method:  "read_abi",
				},
			},
		},
	}

	require.Equal(t, expected, api_err)
}

func TestGetInfoEmptyError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		payload := `{}`
		res.WriteHeader(401)
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	_, err := client.GetHealth(context.Background())
	require.EqualError(t, err, "server returned HTTP 401 Unauthorized")
}

func TestGetHealth(t *testing.T) {
	client := New(testServer.URL)

	h, err := client.GetHealth(context.Background())

	require.NoError(t, err)
	assert.Equal(t, "1.0", h.Version)
	assert.Equal(t, "028d5a34463884fcbe2ecfd3c0fcb3b5d4d538f4fd64803c1ef7209c85f2f266", h.VersionHash)
	assert.Equal(t, 2, len(h.Health))

	assert.Equal(t, "Service1", h.Health[0].Name)
	assert.Equal(t, "OK", h.Health[0].Status)
	assert.Equal(t, time.Time(time.Date(2022, time.January, 14, 15, 39, 41, 678, time.UTC)), h.Health[0].Time)
	assert.Equal(t, 0, len(h.Health[0].Data))

	assert.Equal(t, "Service2", h.Health[1].Name)
	assert.Equal(t, "DOWN", h.Health[1].Status)
	assert.Equal(t, time.Time(time.Date(2022, time.January, 14, 15, 39, 41, 736, time.UTC)), h.Health[1].Time)
	assert.Equal(t, len(h.Health[1].Data), 2)

	assert.Equal(t, float64(1234), h.Health[1].Data["key1"])
	assert.Equal(t, "some_string", h.Health[1].Data["key2"])
}

func TestGetHealthHTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		payload := `{
            "code":404,
            "message":"Not Found",
            "error":{
                "code":0,
                "name":"exception",
                "what":"unspecified",
                "details":[
                    {
                        "message":"Some Error",
                        "file":"file.cpp",
                        "line_number":1337,
                        "method":"some_method"
                    }
                ]
            }
        }`
		res.WriteHeader(404)
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	_, err := client.GetHealth(context.Background())
	require.EqualError(t, err, "404 Not Found")

	api_err, ok := err.(APIError)
	require.True(t, ok)

	expected := APIError{
		Code:    404,
		Message: "Not Found",
		Err: APIErrorInner{
			Code: 0,
			Name: "exception",
			What: "unspecified",
			Details: []APIErrorDetail{
				{
					Message: "Some Error",
					File:    "file.cpp",
					Line:    1337,
					Method:  "some_method",
				},
			},
		},
	}

	require.Equal(t, expected, api_err)
}

func TestGetHealthEmptyError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		payload := `{}`
		res.WriteHeader(500)
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	_, err := client.GetHealth(context.Background())
	require.EqualError(t, err, "server returned HTTP 500 Internal Server Error")
}
