package pkg

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var Collection *mongo.Collection

func GetMongoClient(timeout time.Duration, URI string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}

	/*
		databases, _ := client.ListDatabaseNames(ctx, bson.M{})
		log.Printf("%v", databases)
		collections, _ := client.Database("test").ListCollectionNames(ctx, bson.M{})
		log.Printf("%v", collections)
	*/

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetMongoDbCollection(client *mongo.Client, dbName string, colName string) (*mongo.Collection, error) {
	col := client.Database(dbName).Collection(colName)
	return col, nil
}

