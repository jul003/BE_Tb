package module

import (
	"context"
	"fmt"
	"log"
	"errors"
	"time"
	"github.com/jul003/BE_Tb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


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

func InsertGadget(db *mongo.Database, col string, nama string, merk string, harga float64, spesifikasi model.Spesifikasi, deskripsi string) (insertedID primitive.ObjectID, err error) {
	gadgets := bson.M{
		"nama": nama,
		"merk": merk,
		"harga": harga,
		"spesifikasi": spesifikasi,
		"deskripsi": deskripsi,
	}
	result, err :=db.Collection(col).InsertOne(context.Background(), gadgets)
	if err != nil{
		fmt.Printf("InsertGadget: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}


func GetDataGadget(db *mongo.Database, col string) (data []model.Gadget) {
	gadget := db.Collection(col)
	filter := bson.M{}
	cursor, err := gadget.Find(context.TODO(), filter)
	if err != nil{
		fmt.Println("GetDataGadget: ", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil{
		fmt.Println(err)
	}
	return
}


func InsertReview(db *mongo.Database, col string, rating int, review string ) (insertedID primitive.ObjectID, err error) {
	reviews := bson.M{
		"rating": rating,
		"review": review,
		"datetime": primitive.NewDateTimeFromTime(time.Now().UTC()),
	}
	result, err := db.Collection(col).InsertOne(context.Background(), reviews)
	if err != nil{
		fmt.Printf("InsertReview: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetDataReview(db *mongo.Database, col string) (data []model.Review) {
	reviews := db.Collection(col)
	filter := bson.M{}
	cursor, err := reviews.Find(context.TODO(), filter)
	if err != nil{
		fmt.Println("GetDataReview: ", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil{
		fmt.Println(err)
	}
	return
}

func UpdateGadget(db *mongo.Database, col string, id primitive.ObjectID, nama string, merk string, harga float64, spesifikasi model.Spesifikasi, deskripsi string) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
		"nama": nama,
		"merk": merk,
		"harga": harga,
		"spesifikasi": spesifikasi,
		"deskripsi": deskripsi,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("UpdateGadget: %v", err)
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}
	return nil
}

func DeleteGadgetID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	gadge := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := gadge.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

func GetGadgetByID(_id primitive.ObjectID, db *mongo.Database, col string) (gadgets model.Gadget, errs error) {
	god := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := god.FindOne(context.TODO(), filter).Decode(&gadgets)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return gadgets, fmt.Errorf("no data found for ID %s", _id)
		}
		return gadgets, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return gadgets, nil
}
