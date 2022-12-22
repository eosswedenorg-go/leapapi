package leapapi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInfo_JsonEncode(t *testing.T) {
	info := Info{
		ServerVersion: "a1a94ca",
		HeadBlockNum:  1234,
		HeadBlockTime: time.Date(2020, 4, 20, 6, 33, 22, 0, time.UTC),
	}

	expected := "{\"server_version\":\"a1a94ca\",\"head_block_num\":1234,\"head_block_time\":\"2020-04-20T06:33:22\"}"

	payload, err := json.Marshal(info)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(payload))
}

func TestInfo_JsonDecode(t *testing.T) {
	info := Info{}

	payload := "{\"server_version\":\"94975d6\",\"head_block_num\":888222,\"head_block_time\":\"2019-08-04T13:33:54\"}"

	expected := Info{
		ServerVersion: "94975d6",
		HeadBlockNum:  888222,
		HeadBlockTime: time.Date(2019, 8, 4, 13, 33, 54, 0, time.UTC),
	}

	err := json.Unmarshal([]byte(payload), &info)
	assert.NoError(t, err)
	assert.Equal(t, expected, info)
}
