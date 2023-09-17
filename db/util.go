package db

import "bytes"

func getJoinedKey(table string, key []byte) (newKey []byte) {
	return bytes.Join([][]byte{[]byte(table), key}, []byte("-"))
}

func getSplitTableAndKey(newKey []byte) (table string, key []byte) {
	rs := bytes.Split(newKey, []byte("-"))
	return string(rs[0]), bytes.Join(rs[1:], []byte{})
}
