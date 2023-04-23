package db

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

func GetConnection() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGO_URI")

	clientOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(mongoURI)
		// Connect to MongoDB
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		// Check the connection
		err = client.Ping(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}

		clientInstance = client
	})

	return clientInstance
}

// GetCollection connects to the database and returns the indicated collection
func GetCollection(collection string) *mongo.Collection {
	db := GetConnection()
	mongoDbName := os.Getenv("MONGO_DBNAME")
	return db.Database(mongoDbName).Collection(collection)
}

// FindByUserId searches for a user based on its Id
func FindByUserId(id string) (bson.M, error) {
	userColl := GetCollection("Client")
	var result bson.M
	err := userColl.FindOne(context.Background(), bson.M{"UserId": id}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateUser adds the generated Id to the Client database
func CreateUser(id string) error {
	userColl := GetCollection("Client")
	_, err := userColl.InsertOne(context.Background(), bson.M{"UserId": id})
	if err != nil {
		return err
	}
	return nil

}

// FindByCryptoId searches for a cryptoId from the indicated Id
func FindByCryptoId(id string) (bson.M, error) {
	cryptoColl := GetCollection("Crypto")
	var result bson.M
	err := cryptoColl.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindVoteCrypto searches the user's database if it stores the crypto id
func FindVoteCrypto(idUser string, cryptoId string) (bson.M, error) {
	userColl := GetCollection("Client")
	_, err := FindByUserId(idUser)
	if err != nil {
		return nil, err
	}
	var result bson.M
	err = userColl.FindOne(context.Background(), bson.M{"cryptoId": cryptoId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateVoteValue updates, through the crypto id, the number of votes in the crypto database
func UpdateVoteValue(cryptoId string) error {
	cryptoColl := GetCollection("Crypto")
	_, err := cryptoColl.UpdateOne(context.Background(), bson.M{"_id": cryptoId}, bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "vote", Value: 1}}}}, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

// DecrementVoteValue removes, through the crypto id, a vote from the crypto database
func DecrementVoteValue(cryptoId string) error {
	cryptoColl := GetCollection("Crypto")
	_, err := cryptoColl.UpdateOne(context.Background(), bson.M{"_id": cryptoId}, bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "vote", Value: -1}}}})
	if err != nil {
		return err
	}
	return nil
}

// UpdateVoteValueUser updates, through the crypto id, the number of votes in the user's database
func UpdateVoteValueUser(userId string, cryptoId string) (bson.M, error) {
	userColl := GetCollection("Client")
	var result bson.M
	err := userColl.FindOne(context.Background(), bson.M{"UserId": userId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	_, err = userColl.UpdateOne(context.Background(), bson.M{"UserId": userId}, bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "vote", Value: 1}, primitive.E{Key: "cryptoId", Value: cryptoId}}}}, options.Update().SetUpsert(true))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// DecrementVoteValueUser withdraws, through the crypto id, a vote from the user's database
func DecrementVoteValueUser(userId string, cryptoId string) (bson.M, error) {
	userColl := GetCollection("Client")
	var result bson.M
	err := userColl.FindOne(context.Background(), bson.M{"UserId": userId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	_, err = userColl.UpdateOne(context.Background(), bson.M{"UserId": userId}, bson.D{primitive.E{Key: "$dec", Value: bson.D{primitive.E{Key: "vote", Value: 1}, primitive.E{Key: "cryptoId", Value: cryptoId}}}})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
