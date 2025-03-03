package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbAdapter struct {
	Client       *mongo.Client
	Url          string
	Database     *mongo.Database
	DatabaseName string
}

func NewMongoDbAdapter(url string, databaseName string) *MongoDbAdapter {
	return &MongoDbAdapter{
		Client:       nil,
		Url:          url,
		Database:     nil,
		DatabaseName: databaseName,
	}
}

func (e *MongoDbAdapter) Connect() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(e.Url))
	if err != nil {
		return errors.New("Failed to connect to MongoDB: " + err.Error())
	}

	e.Database = client.Database(e.DatabaseName)

	return nil
}

func (e *MongoDbAdapter) Insert(collectionName string, data map[string]interface{}) error {
	bsonData, err := toBsonD(data)
	if err != nil {
		return errors.New("Failed to convert map to bson.D: " + err.Error())
	}

	collection := e.Database.Collection(collectionName)
	_, err = collection.InsertOne(context.Background(), bsonData)
	if err != nil {
		return errors.New("Failed to insert data: " + err.Error())
	}

	return err
}

func (e *MongoDbAdapter) Disconnect() {
	e.Client.Disconnect(context.Background())
}

func toBsonD(data map[string]interface{}) (bson.D, error) {
	bsonBytes, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}

	var bsonDoc bson.D
	err = bson.Unmarshal(bsonBytes, &bsonDoc)
	if err != nil {
		return nil, err
	}

	return bsonDoc, nil
}
