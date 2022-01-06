package utils

import "go.mongodb.org/mongo-driver/bson"

func ConvertBsonIntoStruct(bsonC interface{}) map[string]interface{}{
	var structR map[string]interface{}
	bsonByte, _ := bson.Marshal(bsonC)
	bson.Unmarshal(bsonByte, &structR)
	return structR
}
