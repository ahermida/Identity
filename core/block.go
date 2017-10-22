package core

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// Header holds a block's header -- we're structuring the JSON
type Header struct {
	Nonce       int64	        `json:"nonce"`
	Difficulty  int64	        `json:"difficulty"`
	Number      int64           `json:"number"`
	Time        int64       	`json:"timestamp"`
	Root        common.Hash     `json:"root"`
	TxHash      common.Hash     `json:"txRoot"`
	ParentHash  common.Hash     `json:"parentHash"`
	Coinbase    common.Address  `json:"miner"`
}

// The in-memory representation of a block
type Block struct {
	header        *Header
	transactions  Transactions
	ReceivedAt    time.Time
	ReceivedFrom  interface{}
}
