package leapapi

import (
	"context"
	"net/url"

	"github.com/imroc/req/v3"
)

type Client struct {
	Url    string
	Host   string
	client *req.Client
}

func New(url string) *Client {
	rc := req.C().
		SetJsonMarshal(json.Marshal).
		SetJsonUnmarshal(customJsonUnmarshal)

	return &Client{
		Url:    url,
		Host:   "",
		client: rc,
	}
}

func (c *Client) send(ctx context.Context, method string, path string, body interface{}, out interface{}) error {
	host := c.Host
	if len(host) < 1 {
		u, err := url.Parse(c.Url)
		if err != nil {
			return err
		}
		host = u.Host
	}

	// Go's net.http (that `req` uses) sends the port in the host header.
	// nodeos api does not like that, so we need to provide our
	// own Host header with just the host.
	r, err := c.client.R().
		SetContext(ctx).
		SetHeader("Host", host).
		SetBody(body).
		Send(method, c.Url+path)

	if err == nil {
		if !r.IsError() {
			err = r.UnmarshalJson(&out)
		} else {
			var api_err APIError
			// Parse error object.
			err = r.UnmarshalJson(&api_err)
			if err != nil || api_err.IsEmpty() {
				// Failed to parse error object. just return an generic HTTP error
				return HTTPError{Code: r.StatusCode}
			}
			err = api_err
		}
	}

	return err
}

//	GetInfo - Fetches "/v1/chain/get_info" from API
//
// ---------------------------------------------------------
func (c *Client) GetInfo(ctx context.Context) (info Info, err error) {
	err = c.send(ctx, "GET", "/v1/chain/get_info", nil, &info)
	return
}

//	Health - Fetches "/v2/health" from API
//
// ---------------------------------------------------------
func (c *Client) GetHealth(ctx context.Context) (health Health, err error) {
	err = c.send(ctx, "GET", "/v2/health", nil, &health)
	return
}
