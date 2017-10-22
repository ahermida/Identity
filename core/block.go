package core

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// Header holds a block's header
// - we're structuring the JSON because we'll be handing this over the web
type Header struct {
	Version 	int 			`json:"version"`
	Nonce       int		        `json:"nonce"`
	Time        int64       	`json:"timestamp"`
	Root        common.Hash     `json:"root"`
	ParentHash  common.Hash     `json:"parentHash"`
}

// The in-memory representation of a block
// - 4 byte block size,
// - transaction counter,
// will be calculated on the fly
type Block struct {
	header        *Header
	transactions  Transactions
}
