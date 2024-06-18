package utils

import (
	"github.com/PhilanderNews/BackendGin/helpers"
	"github.com/PhilanderNews/BackendGin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertVerifikasi(mongoenv *mongo.Database, collname string, dataverifikasi models.Verifikasi) (interface{}, error) {
	return helpers.InsertOneDoc(mongoenv, collname, dataverifikasi)
}

func FindVerifikasi(mconn *mongo.Database, collname string, datauser models.Users) models.Verifikasi {
	filter := bson.M{"username": datauser.Username}
	return helpers.GetOneDoc[models.Verifikasi](mconn, collname, filter)
}

func VerifikasiExists(mconn *mongo.Database, collname string, dataverifikasi models.Verifikasi) bool {
	filter := bson.M{"username": dataverifikasi.Username}
	return helpers.DocExists[models.Verifikasi](mconn, collname, filter, dataverifikasi)
}
