package be63

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func Insertsolana(name string, network string, storage string, ram string, processor string, camera int, display int, battery int, os string, price float64) (InsertedID interface{}) {
	var solana solana
	solana.Name = name
	solana.Network = network
	solana.Storage = storage
	solana.RAM = ram
	solana.Processor = processor
	solana.Camera = camera
	solana.Display = display
	solana.Battery = battery
	solana.OS = os
	solana.Price = price
	return InsertOneDoc("solana", "spec", solana)
}

func GetAllGamme() ([]solana, error) {
	var gem []solana

	cursor, err := MongoConnect("solana").Collection("specs").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error fetching gadgets: %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var gemi solana
		err := cursor.Decode(&gemi)
		if err != nil {
			return nil, fmt.Errorf("error decoding gadget: %v", err)
		}
		gem = append(gem, gemi)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return gem, nil
}

func GetspecbyID(id primitive.ObjectID) (solana, error) {
	var spec solana

	filter := bson.M{"_id": id}
	err := MongoConnect("solana").Collection("spec").FindOne(context.Background(), filter).Decode(&spec)
	if err != nil {
		return solana{}, err
	}

	return spec, nil
}
