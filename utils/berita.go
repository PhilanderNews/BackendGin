package utils

import (
	"github.com/PhilanderNews/BackendGin/helpers"
	"github.com/PhilanderNews/BackendGin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertBerita(mongoenv *mongo.Database, collname string, databerita models.Berita) (interface{}, error) {
	return helpers.InsertOneDoc(mongoenv, collname, databerita)
}

func GetAllBerita(mconn *mongo.Database, collname string) ([]models.Berita, error) {
	return helpers.GetAllDoc[models.Berita](mconn, collname)
}

func FindBerita(mconn *mongo.Database, collname string, databerita models.Berita) models.Berita {
	filter := bson.M{"id": databerita.ID}
	return helpers.GetOneDoc[models.Berita](mconn, collname, filter)
}

func IDBeritaExists(mconn *mongo.Database, collname string, databerita models.Berita) bool {
	filter := bson.M{"id": databerita.ID}
	return helpers.DocExists[models.Berita](mconn, collname, filter, databerita)
}

func UpdateBerita(mconn *mongo.Database, collname string, databerita models.Berita) interface{} {
	filter := bson.M{"id": databerita.ID}
	return helpers.ReplaceOneDoc(mconn, collname, filter, databerita)
}

func DeleteBerita(mconn *mongo.Database, collname string, databerita models.Berita) interface{} {
	filter := bson.M{"id": databerita.ID}
	return helpers.DeleteOneDoc(mconn, collname, filter)
}
