package leapapi

import (
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
		SetJsonUnmarshal(json.Unmarshal)

	return &Client{
		Url:    url,
		Host:   "",
		client: rc,
	}
}

func (c *Client) send(method string, path string) (*req.Response, error) {
	host := c.Host
	if len(host) < 1 {
		u, err := url.Parse(c.Url)
		if err != nil {
			return nil, err
		}
		host = u.Host
	}

	// Go's net.http (that `req` uses) sends the port in the host header.
	// nodeos api does not like that, so we need to provide our
	// own Host header with just the host.
	r, err := c.client.R().
		SetHeader("Host", host).
		Send(method, c.Url+path)

	if err == nil && r.IsError() {
		var api_err APIError
		// Parse error object.
		err = r.UnmarshalJson(&api_err)
		if err != nil || api_err.IsEmpty() {
			// Failed to parse error object. just return an generic HTTP error
			return r, HTTPError{Code: r.StatusCode}
		}
		err = api_err
	}
	return r, err
}

//	GetInfo - Fetches "/v1/chain/get_info" from API
//
// ---------------------------------------------------------
func (c *Client) GetInfo() (Info, error) {
	var info Info

	r, err := c.send("GET", "/v1/chain/get_info")
	if err == nil {
		// Parse json
		err = r.UnmarshalJson(&info)
	}
	return info, err
}

//	Health - Fetches "/v2/health" from API
//
// ---------------------------------------------------------
func (c *Client) GetHealth() (Health, error) {
	var health Health

	r, err := c.send("GET", "/v2/health")
	if err == nil {
		// Parse json
		err = r.UnmarshalJson(&health)
	}
	return health, err
}
