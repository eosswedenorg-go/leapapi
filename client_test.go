
package eosapi

import (
    "time"
    "testing"
    "net/http"
    "net/http/httptest"
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

        res.Write([]byte(info))
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

        res.Write([]byte(info))
    }

}))

func TestGetInfo(t *testing.T) {

    client := New(testServer.URL)

    info, err := client.GetInfo()

    require.NoError(t, err)
    assert.Equal(t, info.ServerVersion, "d1bc8d3")
    assert.Equal(t, info.HeadBlockNum, int64(8888))
    assert.Equal(t, info.HeadBlockTime, time.Unix(1514813821, 0).UTC())
    assert.Equal(t, info.HTTPStatusCode, 200)
}


func TestGetHealth(t *testing.T) {

    client := New(testServer.URL)

    h, err := client.GetHealth()

    require.NoError(t, err)
    assert.Equal(t, h.VersionHash, "028d5a34463884fcbe2ecfd3c0fcb3b5d4d538f4fd64803c1ef7209c85f2f266")
    assert.Equal(t, len(h.Health), 2)

    assert.Equal(t, h.Health[0].Name, "Service1")
    assert.Equal(t, h.Health[0].Status, "OK")
    assert.Equal(t, h.Health[0].Time, int64(1642174781678))
    assert.Equal(t, len(h.Health[0].Data), 0)

    assert.Equal(t, h.Health[1].Name, "Service2")
    assert.Equal(t, h.Health[1].Status, "DOWN")
    assert.Equal(t, h.Health[1].Time, int64(1642174781736))
    assert.Equal(t, len(h.Health[1].Data), 2)

    assert.Equal(t, h.Health[1].Data["key1"], float64(1234))
    assert.Equal(t, h.Health[1].Data["key2"], "some_string")
}
