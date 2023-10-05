package db

import "kloud/cache"

// 这个包是数据库操作的一些语法糖

func setDefaultDB(db DB) {
	defaultDB = db
}

func Get(name string, key []byte) ([]byte, error) {
	//  查缓存，找到直接返回
	c, err := cache.Get(getJoinedKey(name, key))
	if err == nil {
		return c, nil
	}
	return defaultDB.Get(name, key)
}

func Set(name string, key []byte, value []byte) {
	cache.Set(getJoinedKey(name, key), value, 0)
	defaultDB.Set(name, key, value)
}

func Delete(name string, key []byte) error {

	return defaultDB.Delete(name, key)
}

func Range(name string, f func(key []byte, value []byte)) {
	defaultDB.Range(name, f)
}
