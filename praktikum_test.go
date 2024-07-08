package _be2024

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var testClient *mongo.Client

// Set up the MongoDB client for testing
func TestMain(m *testing.M) {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGOSTRING"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	testClient = client

	// Run tests
	code := m.Run()

	// Clean up
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}

	os.Exit(code)
}

func TestInsertGadget(t *testing.T) {
	gadget := Gadget{
		Nama:  "Test Gadget",
		Merk:  "Test Brand",
		Harga: 299.99,
		Spesifikasi: Spesifikasi{
			Prosesor: "Test Prosesor",
			RAM:      8,
			Storage:  256,
			Kamera:   "8 MP",
			Baterai:  "3000 mAh",
			OS:       "Test OS",
			Layar:    "5.5 inch",
			FiturLainnya: []string{"Test Feature 1", "Test Feature 2"},
		},
		Deskripsi: "This is a test gadget.",
	}

	result, err := InsertGadget(testClient, gadget)
	if err != nil {
		t.Fatalf("InsertGadget failed: %v", err)
	}
	if result.InsertedID == nil {
		t.Fatalf("InsertGadget failed: inserted ID is nil")
	}
	t.Logf("Inserted Gadget ID: %v", result.InsertedID)
}

func TestGetGadget(t *testing.T) {
	// First, insert a test gadget to retrieve
	gadget := Gadget{
		Nama:  "Test Gadget",
		Merk:  "Test Brand",
		Harga: 299.99,
		Spesifikasi: Spesifikasi{
			Prosesor: "Test Prosesor",
			RAM:      8,
			Storage:  256,
			Kamera:   "8 MP",
			Baterai:  "3000 mAh",
			OS:       "Test OS",
			Layar:    "5.5 inch",
			FiturLainnya: []string{"Test Feature 1", "Test Feature 2"},
		},
		Deskripsi: "This is a test gadget.",
	}

	insertResult, err := InsertGadget(testClient, gadget)
	if err != nil {
		t.Fatalf("InsertGadget failed: %v", err)
	}

	gadgetID := insertResult.InsertedID.(primitive.ObjectID)
	retrievedGadget, err := GetGadget(testClient, gadgetID)
	if err != nil {
		t.Fatalf("GetGadget failed: %v", err)
	}
	if retrievedGadget.ID != gadgetID {
		t.Fatalf("GetGadget failed: retrieved ID does not match inserted ID")
	}
	t.Logf("Retrieved Gadget: %v", retrievedGadget)
}

func TestInsertReview(t *testing.T) {
	review := Review{
		UserID:   primitive.NewObjectID(),
		GadgetID: primitive.NewObjectID(),
		Rating:   5,
		Review:   "This is a test review.",
		Datetime: primitive.NewDateTimeFromTime(time.Now()),
	}

	result, err := InsertReview(testClient, review)
	if err != nil {
		t.Fatalf("InsertReview failed: %v", err)
	}
	if result.InsertedID == nil {
		t.Fatalf("InsertReview failed: inserted ID is nil")
	}
	t.Logf("Inserted Review ID: %v", result.InsertedID)
}

func TestGetReview(t *testing.T) {
	// First, insert a test review to retrieve
	review := Review{
		UserID:   primitive.NewObjectID(),
		GadgetID: primitive.NewObjectID(),
		Rating:   5,
		Review:   "This is a test review.",
		Datetime: primitive.NewDateTimeFromTime(time.Now()),
	}

	insertResult, err := InsertReview(testClient, review)
	if err != nil {
		t.Fatalf("InsertReview failed: %v", err)
	}

	reviewID := insertResult.InsertedID.(primitive.ObjectID)
	retrievedReview, err := GetReview(testClient, reviewID)
	if err != nil {
		t.Fatalf("GetReview failed: %v", err)
	}
	if retrievedReview.ID != reviewID {
		t.Fatalf("GetReview failed: retrieved ID does not match inserted ID")
	}
	t.Logf("Retrieved Review: %v", retrievedReview)
}
