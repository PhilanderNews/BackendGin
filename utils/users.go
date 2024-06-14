package utils

import (
	"github.com/PhilanderNews/BackendGin/helpers"
	"github.com/PhilanderNews/BackendGin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUser(mongoenv *mongo.Database, collname string, datauser models.Users) (interface{}, error) {
	return helpers.InsertOneDoc(mongoenv, collname, datauser)
}

func GetAllUser(mconn *mongo.Database, collname string) ([]models.Users, error) {
	return helpers.GetAllDoc[models.Users](mconn, collname)
}

func FindUser(mconn *mongo.Database, collname string, userdata models.Users) models.Users {
	filter := bson.M{"username": userdata.Username}
	return helpers.GetOneDoc[models.Users](mconn, collname, filter)
}

func IsPasswordValid(mconn *mongo.Database, collname string, userdata models.Users) bool {
	filter := bson.M{"username": userdata.Username}
	res := helpers.GetOneDoc[models.Users](mconn, collname, filter)
	hashChecker := helpers.CheckPasswordHash(userdata.Password, res.Password)
	return hashChecker
}

func UsernameExists(mconn *mongo.Database, collname string, userdata models.Users) bool {
	filter := bson.M{"username": userdata.Username}
	return helpers.DocExists[models.Users](mconn, collname, filter, userdata)
}

func UpdateUser(mconn *mongo.Database, collname string, datauser models.Users) interface{} {
	filter := bson.M{"username": datauser.Username}
	return helpers.ReplaceOneDoc(mconn, collname, filter, datauser)
}

func DeleteUser(mconn *mongo.Database, collname string, userdata models.Users) interface{} {
	filter := bson.M{"username": userdata.Username}
	return helpers.DeleteOneDoc(mconn, collname, filter)
}
