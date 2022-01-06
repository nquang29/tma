package mongoDB

import (
	"TMA/pkg/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	CollectionRepo
}

//var userCollection = NewCollectionRepo(constants.USER).collection

func (u UserCollection) UserColl() *mongo.Collection{
	return u.NewCollectionRepo(constants.USER).collection
}


func (u UserCollection) Test()  {
	//u.UserColl().FindOne()

}
