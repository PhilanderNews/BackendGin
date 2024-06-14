package utils

import (
	"os"

	"github.com/PhilanderNews/BackendGin/helpers"
	"github.com/PhilanderNews/BackendGin/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(mongoenv, dbname string) *mongo.Database {
	var DBmongoinfo = models.DBInfo{
		DBString: os.Getenv(mongoenv),
		DBName:   dbname,
	}
	return helpers.MongoConnect(DBmongoinfo)
}
