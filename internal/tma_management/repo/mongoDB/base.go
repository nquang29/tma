package mongoDB

import (
	"TMA/internal/database/mongoDB"
	"TMA/pkg/constants"
	"context"
	//	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reader interface {
	FindById(string, string) (interface{}, error)
	FindMulti(interface{}, string) (interface{}, error)
}

type Writer interface {
	Create(interface{}, string) (interface{}, error)
	Update(interface{}, string) (interface{}, error)
}

type BaseCollection interface {
	Reader
	Writer
}

type CollectionRepo struct {
	collection *mongo.Collection
}

var _ BaseCollection = &CollectionRepo{}

var BaseRepo *CollectionRepo

var clientMongo, _ = mongoDB.Init()

func (c *CollectionRepo)NewCollectionRepo(coll string) *CollectionRepo{
	return &CollectionRepo{collection: clientMongo.Database(constants.MONGODB).Collection(coll)}
}

func (c *CollectionRepo) FindById(id string, coll string) (interface{}, error) {
	var result bson.M
	err := c.NewCollectionRepo(coll).collection.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result)
	if err != nil{
		return nil, err
	}
	return result, nil
}

func (c *CollectionRepo) FindMulti(filter interface{}, coll string) (interface{}, error){
	cursor, err := c.NewCollectionRepo(coll).collection.Find(context.TODO(), filter)
	if err != nil{
		return nil, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results, nil
}

func (c CollectionRepo) Create(i interface{}, coll string) (interface{},error) {
	result, err := c.NewCollectionRepo(coll).collection.InsertOne(context.TODO(), i)
	if err != nil{
		return nil, err
	}
	return result, nil
}

func (c CollectionRepo) Update(i interface{}, coll string) (interface{},error) {
	result, err := c.NewCollectionRepo(coll).collection.InsertOne(context.TODO(), i)
	if err != nil{
		return nil, err
	}
	return result, nil
}


