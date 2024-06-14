package helpers

import (
	"context"
	"fmt"

	"github.com/PhilanderNews/BackendGin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(mconn models.DBInfo) (db *mongo.Database) {
	clientOptions := options.Client().ApplyURI((mconn.DBString))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v", err)
	}
	return client.Database(mconn.DBName)
}

func Create2dsphere(mconn models.DBInfo) (db *mongo.Database) {
	clientOptions := options.Client().ApplyURI((mconn.DBString))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v", err)
	}

	// Mengecek apakah index sudah ada
	collection := client.Database(mconn.DBName).Collection(mconn.CollectionName)
	cursor, err := collection.Indexes().List(context.TODO())
	if err != nil {
		fmt.Printf("Error listing indexes: %v", err)
	}

	indexExists := false
	for cursor.Next(context.TODO()) {
		var index bson.M
		if err := cursor.Decode(&index); err != nil {
			fmt.Printf("Error decoding index: %v", err)
		}
		if index["name"] == "geometry_2dsphere" {
			indexExists = true
			break
		}
	}

	// Membuat indeks jika belum ada
	if !indexExists {
		indexModel := mongo.IndexModel{
			Keys: bson.D{
				{Key: "geometry", Value: "2dsphere"},
			},
		}

		_, err = client.Database(mconn.DBName).Collection(mconn.CollectionName).Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			fmt.Printf("Error creating geospatial index: %v", err)
		}
	}
	return client.Database(mconn.DBName)
}

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}, err error) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
		return nil, err
	}
	return insertResult.InsertedID, nil
}

func GetOneDoc[T any](db *mongo.Database, collection string, filter bson.M) (doc T) {
	err := db.Collection(collection).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("GetOneDoc: %v\n", err)
	}
	return
}

func GetOneLatestDoc[T any](db *mongo.Database, collection string, filter bson.M) (doc T, err error) {
	opts := options.FindOne().SetSort(bson.M{"$natural": -1})
	err = db.Collection(collection).FindOne(context.TODO(), filter, opts).Decode(&doc)
	return doc, err
}

func GetAllDocByFilter[T any](db *mongo.Database, collection string, filter bson.M) (doc []T, err error) {
	ctx := context.TODO()
	cur, err := db.Collection(collection).Find(ctx, filter)
	if err != nil {
		fmt.Printf("GetAllDoc: %v\n", err)
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &doc)
	if err != nil {
		fmt.Printf("GetAllDoc Cursor Err: %v\n", err)
		return nil, err
	}
	return doc, nil
}

func GetAllDoc[T any](db *mongo.Database, collection string) (doc []T, err error) {
	ctx := context.TODO()
	cur, err := db.Collection(collection).Find(ctx, bson.M{})
	if err != nil {
		fmt.Printf("GetAllDoc: %v\n", err)
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &doc)
	if err != nil {
		fmt.Printf("GetAllDoc Cursor Err: %v\n", err)
		return nil, err
	}
	return doc, nil
}

func GetAllDistinctDoc(db *mongo.Database, filter bson.M, fieldname, collection string) (doc []any) {
	ctx := context.TODO()
	doc, err := db.Collection(collection).Distinct(ctx, fieldname, filter)
	if err != nil {
		fmt.Printf("GetAllDistinctDoc: %v\n", err)
	}
	return
}

func ReplaceOneDoc(db *mongo.Database, collection string, filter bson.M, doc interface{}) (updatereseult *mongo.UpdateResult) {
	updatereseult, err := db.Collection(collection).ReplaceOne(context.TODO(), filter, doc)
	if err != nil {
		fmt.Printf("ReplaceOneDoc: %v\n", err)
	}
	return
}

func DeleteOneDoc(db *mongo.Database, collection string, filter bson.M) (result *mongo.DeleteResult) {
	result, err := db.Collection(collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteOneDoc: %v\n", err)
	}
	return
}

func DeleteDoc(db *mongo.Database, collection string, filter bson.M) (result *mongo.DeleteResult) {
	result, err := db.Collection(collection).DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDoc : %v\n", err)
	}
	return
}

func DocExists[T any](db *mongo.Database, collname string, filter bson.M, doc T) (result bool) {
	err := db.Collection(collname).FindOne(context.Background(), filter).Decode(&doc)
	return err == nil
}
