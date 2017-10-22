/*
        Package core contains most of the blockchain logic
*/
package core

import (
        "fmt"
        "log"
        "sync"

        "github.com/ahermida/Identity/db"
        "github.com/ahermida/Identity/config"
        "github.com/hashicorp/golang-lru"
)

type BlockChain struct {
	config         config.Network   // network configuration
    params         config.Params    // chain configuration

	chainDB        db.Database      // DB interface

	genesisBlock   *Block           // First block
    tip            *Block           // Current head of the block chain

	mu             sync.RWMutex     // lock all chain ops
	chainmu        sync.RWMutex     // blockchain insertion lock
	procmu         sync.RWMutex     // block processor lock

    blockCache     *lru.Cache       // Cache for the most recent entire blocks
	futureBlocks   *lru.Cache       // future blocks are blocks added for later processing
	wg             sync.WaitGroup   // chain processing wait group for shutting down

	engine         ConsensusEngine  // Does PoW for us
	processor      Processor        // Processes new blocks for us
	validator      Validator        // block and state validator interface
}
