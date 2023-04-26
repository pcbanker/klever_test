package db

import (
	"context"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// TestGetConnection is a unit test function that tests the GetConnection function, which establishes a connection to a MongoDB database.
func TestGetConnection(t *testing.T) {

	err := godotenv.Load("./../.env") //Loads environment variables from the .env file.
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := GetConnection() //Calls the GetConnection function to obtain a client instance.

	if client == nil {
		t.Error("Expected client instance, got nil") //Checks that the client instance is nil.
	}

	err = client.Ping(context.Background(), nil) //Pings the MongoDB server to check if the connection was established successfully.
	if err != nil {
		t.Errorf("Failed to establish connection to MongoDB: %v", err)
	}

	client2 := GetConnection() //Calls the GetConnection function again to obtain another client instance.

	assert.Equal(t, client, client2) //Uses the assert.Equal function to check that the two client instances are equal.
}

// TestGetCollection is a unit test function that tests the GetCollection function, which retrieves a MongoDB collection based on the provided name.
func TestGetCollection(t *testing.T) {

	err := godotenv.Load("./../.env") //Loads environment variables from the .env file.
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	collection := GetCollection("Crypto") //Calls the GetCollection function with the collection name "Crypto".

	assert.NotEmpty(t, collection) //Checks that the collection returned is not empty
}
