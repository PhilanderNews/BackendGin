package utils

import (
	"github.com/PhilanderNews/BackendGin/helpers"
	"github.com/PhilanderNews/BackendGin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertKomentar(mongoenv *mongo.Database, collname string, datakomentar models.Komentar) (interface{}, error) {
	return helpers.InsertOneDoc(mongoenv, collname, datakomentar)
}

func GetAllKomentar(mconn *mongo.Database, collname string) ([]models.Komentar, error) {
	return helpers.GetAllDoc[models.Komentar](mconn, collname)
}

func FindKomentar(mconn *mongo.Database, collname string, datakomentar models.Komentar) models.Komentar {
	filter := bson.M{"id": datakomentar.ID}
	return helpers.GetOneDoc[models.Komentar](mconn, collname, filter)
}

func IDKomentarExists(mconn *mongo.Database, collname string, datakomentar models.Komentar) bool {
	filter := bson.M{"id": datakomentar.ID}
	return helpers.DocExists[models.Komentar](mconn, collname, filter, datakomentar)
}

func UpdateKomentar(mconn *mongo.Database, collname string, datakomentar models.Komentar) interface{} {
	filter := bson.M{"id": datakomentar.ID}
	return helpers.ReplaceOneDoc(mconn, collname, filter, datakomentar)
}

func DeleteKomentar(mconn *mongo.Database, collname string, datakomentar models.Komentar) interface{} {
	filter := bson.M{"id": datakomentar.ID}
	return helpers.DeleteOneDoc(mconn, collname, filter)
}
