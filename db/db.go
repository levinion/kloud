package db

// 数据库接口。数据库定义：表（Table）的集合
type DB interface {
	Init(table string)                                    //初始化表
	Get(table string, key []byte) ([]byte, error)         //从表中取值
	Set(table string, key []byte, value []byte)           // 设置表中的字段
	Delete(table string, key []byte) error                // 删除表中字段
	Range(table string, f func(key []byte, value []byte)) //遍历表中的字段
}

var defaultDB DB = nil
