package db

import (
	"kloud/constant"
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger"
)

type BadgerDB struct {
	Tables map[string]*badger.DB
}

func UseBadgerDB() DB {
	db := &BadgerDB{make(map[string]*badger.DB)}
	setDefaultDB(db)
	return db
}

// 惰性初始化，用到时自动打开
func (db *BadgerDB) Init(name string) {
	path := filepath.Join(constant.DBPath, "badger", name)
	opt := badger.DefaultOptions(path)
	opt.Logger = nil
	os.MkdirAll(path, os.ModePerm)
	instance, err := badger.Open(opt)
	if err != nil {
		panic(err)
	}
	db.Tables[name] = instance
}

func (db *BadgerDB) Get(name string, key []byte) ([]byte, error) {
	var value []byte
	table, ok := db.Tables[name]
	if !ok {
		db.Init(name)
		table = db.Tables[name]
	}
	err := table.View(func(txn *badger.Txn) error {
		ob, err := txn.Get(key)
		if err != nil {
			return err
		}
		ob.Value(func(val []byte) error {
			value = val
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (db *BadgerDB) Set(name string, key []byte, value []byte) {

	table, ok := db.Tables[name]
	if !ok {
		db.Init(name)
		table = db.Tables[name]
	}
	//You can't get error!
	table.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

func (db *BadgerDB) Delete(name string, key []byte) error {
	table, ok := db.Tables[name]
	if !ok {
		db.Init(name)
		table = db.Tables[name]
	}
	return table.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func (db *BadgerDB) Range(name string, f func(key []byte, value []byte)) {
	table, ok := db.Tables[name]
	if !ok {
		db.Init(name)
		table = db.Tables[name]
	}
	table.View(func(txn *badger.Txn) error {
		iter := txn.NewIterator(badger.DefaultIteratorOptions)
		for iter.Rewind(); iter.Valid(); iter.Next() {
			item := iter.Item()
			item.Value(func(val []byte) error {
				f(item.Key(), val)
				return nil
			})
		}
		iter.Close()
		return nil
	})
}
