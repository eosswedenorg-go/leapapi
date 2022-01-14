
package eosapi;

import (
    "strings"
    "time"
    "net/url"
    "io/ioutil"
    "github.com/imroc/req"
    jsontime "github.com/liamylian/jsontime/v2/v2"
)

var json = jsontime.ConfigWithCustomTimeFormat

func init() {

    // EOS Api does not specify timezone in timestamps (they are always UTC tho).
    jsontime.SetDefaultTimeFormat("2006-01-02T15:04:05", time.UTC)
}

func New(url string) (*Client) {
    return &Client{
        Url: url,
        Host: "",
    }
}

func (c *Client) send(method string, path string) (*req.Resp, error) {

    host := c.Host
    if len(host) < 1 {
        u, err := url.Parse(c.Url)
        if err != nil {
            return nil, err
        }
        host = strings.Split(u.Host, ":")[0]
    }

    // Go's net.http (that `req` uses) sends the port in the host header.
    // nodeos api does not like that, so we need to provide our
    // own Host header with just the host.
    headers := req.Header{
        "Host": host,
    }

    r := req.New()
    return r.Do(method, c.Url + path, headers)
}

//  GetInfo - Fetches "/v1/chain/get_info" from API
// ---------------------------------------------------------
func (c *Client) GetInfo() (Info, error) {

    var info Info

    r, err := c.send("GET", "/v1/chain/get_info")
    if err == nil {
        resp := r.Response()
        body, _ := ioutil.ReadAll(resp.Body)

        // Set HTTPStatusCode
        info.HTTPStatusCode = resp.StatusCode

        // Parse json
        err = json.Unmarshal(body, &info)
    }
    return info, err
}

//  Health - Fetches "/v2/health" from API
// ---------------------------------------------------------
func (c *Client) GetHealth() (Health, error) {

    var health Health;

    r, err := c.send("GET", "/v2/health")
    if err == nil {
        resp := r.Response()
        body, _ := ioutil.ReadAll(resp.Body)

        // Set HTTPStatusCode
        health.HTTPStatusCode = resp.StatusCode

        // Parse json
        err = json.Unmarshal(body, &health)
    }
    return health, err
}
