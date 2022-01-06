package service

import (
	"TMA/internal/tma_management/repo/mongoDB"
	"TMA/pkg/constants"
	"TMA/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Create(user models.User) (models.User, error){
	userCreated, err := mongoDB.BaseRepo.Create(user, constants.USER)
	if err != nil {
		return models.User{}, err
	}
	var result models.User
	bsonByte, _ := bson.Marshal(userCreated)
	bson.Unmarshal(bsonByte, &user)
	return result, nil
}

func Find(id string) (models.User, error){
	userFound, err := mongoDB.BaseRepo.FindById(id, constants.USER)
	if err != nil {
		return models.User{}, err
	}
	//convert bson.M into model
	var user models.User
	bsonByte, _ := bson.Marshal(userFound)
	bson.Unmarshal(bsonByte, &user)
	//user1 := utils.ConvertBsonIntoStruct(userFound)
	//mapstructure.Decode(user1,&user)
	return user, nil
}
