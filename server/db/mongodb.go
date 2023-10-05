package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	database *mongo.Database
	Tables   map[string]*mongo.Collection
}

func UseMongoDB() DB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	db := &MongoDB{client.Database("default"), make(map[string]*mongo.Collection)}
	setDefaultDB(db)
	return db
}

// 惰性初始化，用到时自动打开
func (db *MongoDB) Init(name string) {
	db.Tables[name] = db.database.Collection(name)
}

func (db *MongoDB) Get(name string, key []byte) ([]byte, error) {
	table, ok := db.Tables[name]
	if !ok {
		db.Init(name)
		table = db.Tables[name]
	}
	result := new(bson.E)
	table.FindOne(context.TODO(), bson.E{
		Key: string(key),
	}).Decode(result)
	return result.Value.([]byte), nil
}

func (db *MongoDB) Set(name string, key []byte, value []byte) {

	table, ok := db.Tables[name]
	if !ok {
		db.Init(name)
		table = db.Tables[name]
	}
	//You can't get error!
	table.InsertOne(context.Background(), bson.E{
		Key: string(key), Value: value,
	})
}

func (db *MongoDB) Delete(name string, key []byte) error {
	table, ok := db.Tables[name]
	if !ok {
		db.Init(name)
		table = db.Tables[name]
	}
	_, err := table.DeleteOne(context.TODO(), bson.E{
		Key: string(key),
	})
	return err
}

func (db *MongoDB) Range(name string, f func(key []byte, value []byte)) {
	table, ok := db.Tables[name]
	if !ok {
		db.Init(name)
		table = db.Tables[name]
	}
	cur, _ := table.Find(context.TODO(), bson.E{})
	for cur.Next(context.TODO()) {
		r := new(bson.E)
		cur.Decode(r)
		f([]byte(r.Key), r.Value.([]byte))
	}
}
