package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

// This function is responsible for getting a connection to the MongoDB server and returning it in the form of a pointer (*mongo.Client) so that it can be used in other parts of the code.
func GetConnection() *mongo.Client {

	mongoURI := os.Getenv("MONGO_URI") // mongoURI gets the value from the environment variable and stores it in the variable

	//Executes an anonymous function only once, using the global variable "clientOnce"
	clientOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(mongoURI) //Creates a client options object for MongoDB using the value of the "mongoURI" variable.

		client, err := mongo.Connect(context.Background(), clientOptions) //Connects to MongoDB server using defined client options
		if err != nil {
			log.Fatal(err) //If an error occurred, the function returns it.
		}

		err = client.Ping(context.Background(), nil) //Tests the client's connection to the server
		if err != nil {
			log.Fatal(err) //If an error occurred, the function returns it.
		}

		clientInstance = client //Sets the global variable "clientInstance" to the client object.
	})

	return clientInstance

}

// This function returns a pointer to a MongoDB collection specified as an argument.
// GetCollection takes a string collection as a parameter and returns a pointer to mongo.Collection.
func GetCollection(collection string) *mongo.Collection {

	db := GetConnection() //db gets the GetConnection function to get a MongoDB database connection

	mongoDbName := os.Getenv("MONGO_DBNAME") //mongoDbName get database name value from environment variable

	return db.Database(mongoDbName).Collection(collection) // Returns the specified collection using the database name and collection name
}

// This function increments the value of the "vote" field from one of the "Crypto" collection in the database.
// UpdateVoteValue takes a cryptoId string as a parameter
func UpdateVoteValue(cryptoId string) error {

	cryptoColl := GetCollection("Crypto") //cryptoColl gets a connection to the database "Crypto" collection.

	_, err := cryptoColl.UpdateOne(context.Background(), bson.M{"_id": cryptoId}, bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "vote", Value: 1}}}}) //UpdateOne serves to update a document in the "Crypto" collection. First, the context of the operation is passed, as a second parameter, a filter that searches for a crypto with an id equal to cryptoId, and as a third parameter, an operator that increments the value of the "vote" field by 1.
	if err != nil {
		fmt.Println(err) // If there are any errors, the function returns them.
		return err
	}
	return nil
}

// This function decrements the value of the "vote" field from one of the "Crypto" collection in the database.
// DecrementVoteValue takes a cryptoId string as a parameter
func DecrementVoteValue(cryptoId string) error {

	cryptoColl := GetCollection("Crypto") //cryptoColl gets a connection to the database "Crypto" collection.

	_, err := cryptoColl.UpdateOne(context.Background(), bson.M{"_id": cryptoId}, bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "vote", Value: -1}}}}) //UpdateOne serves to update a document in the "Crypto" collection. First, the context of the operation is passed, as a second parameter, a filter that searches for a crypto with the id equal to cryptoId, and as a third parameter, an operator that increments the value of the "vote" field by -1, that is, decreases by 1.
	if err != nil {
		return err // If there are any errors, the function returns them.
	}
	return nil
}
