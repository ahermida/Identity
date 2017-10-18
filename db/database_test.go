/*
		Test out our CRUD ops with LevelDB
*/
package db_test

import (
		"bytes"
		"io/ioutil"
		"os"
		"testing"

		"github.com/ahermida/Identity/db"
)

func newTestDB() (*db.LDB, func()) {
	dirname, err := ioutil.TempDir(os.TempDir(), "db_test")
	if err != nil {
		panic("Couldn't make test DB: " + err.Error())
	}
	ldb, err := db.NewDB(dirname)
	if err != nil {
		panic("failed to create test database: " + err.Error())
	}

	//pass callback that removes test database from memory
	return ldb, func() {
		ldb.Close()
		os.RemoveAll(dirname)
	}
}

//test our new CRUD funcs
func TestDB_CRUD(t *testing.T) {
	db, rmdb := newTestDB() //generate our test DB
	defer rmdb() 						//delete our DB after test
	testCRUD(db, t) 				//run test
}

//test CRUD functionality
func testCRUD(ldb db.Database, t *testing.T) {

	data := []string{"d", "di", "dingo is my lingo", "8675309"}

	//add data
	for _, val := range data {
		err := ldb.Put([]byte(val), []byte(val))
		if err != nil {
			t.Fatalf("put failed %v", err)
		}
	}

	//get data
	for _, val := range data {
		str, err := ldb.Get([]byte(val))
		if err != nil {
			t.Fatalf("get failed %v", err)
		}
		if !bytes.Equal(str, []byte(val)) {
			t.Fatalf("get returned wrong result, got %q expected %q", string(str), val)
		}
	}

	//update some data
	for _, val := range data {
		err := ldb.Put([]byte(val), []byte("yoyo"))
		if err != nil {
			t.Fatalf("update failed %v", err)
		}
	}

	//check if our update worked
	for _, val := range data {
		str, err := ldb.Get([]byte(val))
		if err != nil {
			t.Fatalf("get failed: %v", err)
		}
		if !bytes.Equal(str, []byte("yoyo")) {
			t.Fatalf("got %q, expected yoyo", string(str))
		}
	}

	//delete all the data we created
	for _, val := range data {
		err := ldb.Delete([]byte(val))
		if err != nil {
			t.Fatalf("Delete %q failed: %v", val, err)
		}
	}

	//check if it all got deleted
	for _, val := range data {
		_, err := ldb.Get([]byte(val))
		if err == nil {
			t.Fatalf("Was able to get a value that should've been deleted", val)
		}
	}
}
