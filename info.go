package leapapi

import "time"

// get_info format
type Info struct {
	HTTPStatusCode int `json:"-"`

	ServerVersion             string    `json:"server_version"`
	ServerVersionString       string    `json:"server_version_string"`
	ServerFullVersionString   string    `json:"server_full_version_string"`
	ChainID                   string    `json:"chain_id"`
	HeadBlockID               string    `json:"head_block_id"`
	HeadBlockNum              int64     `json:"head_block_num"`
	HeadBlockTime             time.Time `json:"head_block_time"`
	HeadBlockProducer         string    `json:"head_block_producer"`
	LastIrreversableBlockNum  int64     `json:"last_irreversible_block_num"`
	LastIrreversableBlockID   string    `json:"last_irreversible_block_id"`
	LastIrreversableBlockTime time.Time `json:"last_irreversible_block_time,omitempty"`
	VirtualBlockCPULimit      int64     `json:"virtual_block_cpu_limit"`
	VirtualBlockNETLimit      int64     `json:"virtual_block_net_limit"`
	BlockCPULimit             int64     `json:"block_cpu_limit"`
	BlockNETLimit             int64     `json:"block_net_limit"`
	TotalCPUWeight            int64     `json:"total_cpu_weigth,omitempty"`
	TotalNETWeight            int64     `json:"total_net_weigth,omitempty"`
	ForkDBHeadBlockID         string    `json:"fork_db_head_block_id"`
	ForkDBHeadBlockNum        int64     `json:"fork_db_head_block_num"`
	EarliestAvailableBlockNum int64     `json:"earliest_available_block_num,omitempty"`
}
