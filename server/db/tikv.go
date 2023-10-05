package db

import (
	"context"

	"github.com/tikv/client-go/v2/txnkv"
)

type TiKV struct {
	client *txnkv.Client
}

func UseTiKV() DB {
	client, err := txnkv.NewClient([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}
	db := &TiKV{client}
	setDefaultDB(db)
	return db
}

func (db *TiKV) Init(name string) {
}

func (db *TiKV) Get(name string, key []byte) ([]byte, error) {
	txn, err := db.client.Begin()
	if err != nil {
		return nil, err
	}
	return txn.Get(context.TODO(), getJoinedKey(name, key))
}

func (db *TiKV) Set(name string, key []byte, value []byte) {
	txn, err := db.client.Begin()
	if err != nil {
		panic(err)
	}
	txn.Set(getJoinedKey(name, key), value)
}

func (db *TiKV) Delete(name string, key []byte) error {
	txn, err := db.client.Begin()
	if err != nil {
		return err
	}
	return txn.Delete(getJoinedKey(name, key))
}

func (db *TiKV) Range(name string, f func(key []byte, value []byte)) {
	txn, _ := db.client.Begin()
	iter, _ := txn.Iter([]byte(""), nil)
	defer iter.Close()
	for iter.Valid() {
		_, key := getSplitTableAndKey(iter.Key())
		f(key, iter.Value())
	}
}
