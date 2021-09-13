package cache

import (

	// "json"

	// "time"

	"github.com/dgraph-io/badger/v3"
)

// var DB *badger.DB
type BadgerDB struct {
	DB *badger.DB
}

func Connect() (*BadgerDB, error) {
	// opts := badger.DefaultOptions("./")
	// opts.Dir = "./badger"
	// opts.ValueDir = "./badger"
	// db, err := badger.Open(opts)
	// if err != nil {
	// 	return &BadgerDB{}, err
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// for i := 1; i < 1000; i++ {
	// 	time.Sleep(1 * time.Second)
	// 	log.Println(i)
	// }
	// return &BadgerDB{
	// 	DB: db,
	// }, nil
	// Ope the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	DB, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	// defer DB.Close()
	if err != nil {
		return &BadgerDB{}, err
	}
	return &BadgerDB{
		DB: DB,
	}, nil
}

func (Bdb *BadgerDB) Set(key string, val []byte) error {
	err := Bdb.DB.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), val)
	})

	return err
}

func (Bdb *BadgerDB) Get(key string) ([]byte, error) {
	var val []byte
	err := Bdb.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		// Use item here.
		val, err = item.ValueCopy(val)
		return err
	})
	if err != nil {
		return val, err
	}
	return val, err
}

func (Bdb *BadgerDB) Delete(key string) error {
	err := Bdb.DB.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})
	return err
}
