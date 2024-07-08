package _be2024

import (
	"context"
	"time"
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

func InsertGadget(client *mongo.Client, gadget Gadget) (*mongo.InsertOneResult, error) {
	collection := client.Database("tb2024").Collection("gadget2024")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, gadget)
	return result, err
}


func GetGadget(client *mongo.Client, id primitive.ObjectID) (Gadget, error) {
	collection := client.Database("tb2024").Collection("gadget2024")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var gadget Gadget
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&gadget)
	return gadget, err
}


func InsertReview(client *mongo.Client, review Review) (*mongo.InsertOneResult, error) {
	collection := client.Database("tb2024").Collection("review2024")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, review)
	return result, err
}


func GetReview(client *mongo.Client, id primitive.ObjectID) (Review, error) {
	collection := client.Database("tb2024").Collection("review2024")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var review Review
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&review)
	return review, err
}

