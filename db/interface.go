/*
		Interfaces for LevelDB
*/
package db

// Database is a wrapper for db ops
type Database interface {
	Put(key []byte, value []byte) error
	Has(key []byte) (bool, error)
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	Close()
	NewBatch() Batch
}

type Batch interface {
	Put(key []byte, value []byte) error
	Write() error
}
