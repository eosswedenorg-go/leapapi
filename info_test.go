package leapapi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInfo_JsonEncode(t *testing.T) {
	info := Info{
		ServerVersion:            "c1c8ed71",
		ServerVersionString:      "v2.0.9-1.20.2",
		ServerFullVersionString:  "v2.0.9-1.20.2-c1c8ed71bc6369f84de706e3362a42db13c06590",
		ChainID:                  "a9c481dfbc7d9506dc7e87e9a137c931b0a9303f64fd7a1d08b8230133920097",
		HeadBlockID:              "05ae4e7033dc0138bb3bede566d57cc7783ba1a091388c6961e93c0ef37476dc",
		HeadBlockNum:             95309424,
		HeadBlockTime:            time.Date(2022, 12, 22, 13, 56, 3, 0, time.UTC),
		HeadBlockProducer:        "eosriobrazil",
		LastIrreversableBlockNum: 95309332,
		LastIrreversableBlockID:  "05ae4e14d2c8d65eb6364cf65f588761f8e8156f9a2ef6b9f53fab1609066f9c",
		VirtualBlockCPULimit:     400000,
		VirtualBlockNETLimit:     1048576,
		BlockCPULimit:            399046,
		BlockNETLimit:            1046064,
		ForkDBHeadBlockNum:       95309424,
		ForkDBHeadBlockID:        "05ae4e7033dc0138bb3bede566d57cc7783ba1a091388c6961e93c0ef37476dc",
	}

	expected := `{
  "server_version": "c1c8ed71",
  "server_version_string": "v2.0.9-1.20.2",
  "server_full_version_string": "v2.0.9-1.20.2-c1c8ed71bc6369f84de706e3362a42db13c06590",
  "chain_id": "a9c481dfbc7d9506dc7e87e9a137c931b0a9303f64fd7a1d08b8230133920097",
  "head_block_id": "05ae4e7033dc0138bb3bede566d57cc7783ba1a091388c6961e93c0ef37476dc",
  "head_block_num": 95309424,
  "head_block_time": "2022-12-22T13:56:03",
  "head_block_producer": "eosriobrazil",
  "last_irreversible_block_num": 95309332,
  "last_irreversible_block_id": "05ae4e14d2c8d65eb6364cf65f588761f8e8156f9a2ef6b9f53fab1609066f9c",
  "virtual_block_cpu_limit": 400000,
  "virtual_block_net_limit": 1048576,
  "block_cpu_limit": 399046,
  "block_net_limit": 1046064,
  "fork_db_head_block_id": "05ae4e7033dc0138bb3bede566d57cc7783ba1a091388c6961e93c0ef37476dc",
  "fork_db_head_block_num": 95309424
}`

	payload, err := json.MarshalIndent(info, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, expected, string(payload))
}

func TestInfo_JsonDecode(t *testing.T) {
	info := Info{}

	// payload := "{\"server_version\":\"94975d6\",\"head_block_num\":888222,\"head_block_time\":\"2019-08-04T13:33:54\"}"
	payload := `{
		"server_version": "94975d6",
		"server_version_string": "v2.0.0",
		"server_full_version_string": "v2.0.0-c1c8ed71bc6369f84de706e3362a42db13c06590",
		"chain_id": "999a4c322ad2891c482dc7c08044442a687e75b2b8d423e3220419ca008b49a8",
		"head_block_id": "1de4ff2f740f581aa2451d2b62d23309ee039941c3d98ba79297d1d5c5b18822",
		"head_block_num": 236718321,
		"head_block_time": "2019-08-04T13:33:54",
		"head_block_producer": "cryptoking",
		"last_irreversible_block_num": 1287389127381,
		"last_irreversible_block_id": "01ae227766a85425bc359da93975bccb8a472a6b9937fce75031840c654ce771",
		"virtual_block_cpu_limit": 600000,
		"virtual_block_net_limit": 10023782,
		"block_cpu_limit": 23817312,
		"block_net_limit": 199202322,
		"fork_db_head_block_num": 21783912781,
		"fork_db_head_block_id": "049995effdfef39ba593603d4e1befbacd113520f55cf00afebf9d25ae336c21"
	  }`

	expected := Info{
		ServerVersion:            "94975d6",
		ServerVersionString:      "v2.0.0",
		ServerFullVersionString:  "v2.0.0-c1c8ed71bc6369f84de706e3362a42db13c06590",
		ChainID:                  "999a4c322ad2891c482dc7c08044442a687e75b2b8d423e3220419ca008b49a8",
		HeadBlockID:              "1de4ff2f740f581aa2451d2b62d23309ee039941c3d98ba79297d1d5c5b18822",
		HeadBlockNum:             236718321,
		HeadBlockTime:            time.Date(2019, 8, 4, 13, 33, 54, 0, time.UTC),
		HeadBlockProducer:        "cryptoking",
		LastIrreversableBlockNum: 1287389127381,
		LastIrreversableBlockID:  "01ae227766a85425bc359da93975bccb8a472a6b9937fce75031840c654ce771",
		VirtualBlockCPULimit:     600000,
		VirtualBlockNETLimit:     10023782,
		BlockCPULimit:            23817312,
		BlockNETLimit:            199202322,
		ForkDBHeadBlockID:        "049995effdfef39ba593603d4e1befbacd113520f55cf00afebf9d25ae336c21",
		ForkDBHeadBlockNum:       21783912781,
	}

	err := json.Unmarshal([]byte(payload), &info)
	assert.NoError(t, err)
	assert.Equal(t, expected, info)
}
