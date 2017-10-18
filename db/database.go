/*
        Package db adds a layer of abstraction to levelDB
        LevelDB API available for reference at:
        https://godoc.org/github.com/syndtr/goleveldb/leveldb
*/
package db

import (
        log "github.com/ahermida/Identity/log"
        "github.com/syndtr/goleveldb/leveldb"
        "github.com/syndtr/goleveldb/leveldb/iterator"
        "github.com/syndtr/goleveldb/leveldb/errors"
      	"github.com/syndtr/goleveldb/leveldb/opt"
        "github.com/syndtr/goleveldb/leveldb/filter"
)

//LevelDB Wrapping Struct
type LDB struct {
    db *leveldb.DB //reference to leveldb instance
    path string    //path to db file
    log log.Logger //logger for reporting errors
}

//Returns a wrapped DB object
func NewDB(filePath string) (*LDB, error) {
    ldb, err := leveldb.OpenFile(filePath, &opt.Options{
        Filter: filter.NewBloomFilter(10),
    })
    if _, corrupt := err.(*errors.ErrCorrupted); corrupt {
        ldb, err = leveldb.RecoverFile(filePath, nil)
    }

    logger := log.New("db", filePath)

    //catch errors if db if we don't have a db
    if err != nil {
        return nil, err
    }

    //return new wrapped db
    return &LDB {
        db: ldb,
        path: filePath,
        log: logger,
    }, nil
}

//Gets location of DB file
func (db *LDB) Location() string {
    return db.path
}

//Puts data into the database
func (db *LDB) Put(key []byte, value []byte) error {
    return db.db.Put(key, value, nil)
}

//Checks for key existence
func (db *LDB) Has(key []byte) (bool, error) {
    return db.db.Has(key, nil)
}

//Gets value for given key, or nil
func (db *LDB) Get(key []byte) ([]byte, error) {
    dat, err := db.db.Get(key, nil)
    if err != nil {
        return nil, err
    }
    return dat, nil
}

//Deletes key/value pair
func (db *LDB) Delete(key []byte) error {
    return db.db.Delete(key, nil)
}

//Gets LDB Iterator
func (db *LDB) NewIterator() iterator.Iterator {
    return db.db.NewIterator(nil, nil)
}

//Closes DB
func (db *LDB) Close() {
    err := db.db.Close()
    if err == nil {
        db.log.Info("LevelDB closed")
    } else {
        db.log.Error("Couldn't close database", "err", err)
    }
}

//Get LDB instance
func (db *LDB) GetDB() *leveldb.DB {
    return db.db
}

//Wrapper for LDB batch edits
type LDBBatch struct {
    db *leveldb.DB
    batch *leveldb.Batch
}

//Returns a new batch object
func (db *LDB) NewBatch() Batch {
    return &LDBBatch{
        db: db.db,
        batch: new(leveldb.Batch),
    }
}

//Adds a new Put to the batch
func (b *LDBBatch) Put(key []byte, value []byte) error {
    b.batch.Put(key, value)
    return nil
}

//Submits batch to DB
func (b *LDBBatch) Write() error {
    return b.db.Write(b.batch, nil)
}
