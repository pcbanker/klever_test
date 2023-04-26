package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pcbanker/test/model"
	"github.com/pcbanker/test/price"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/pcbanker/test/db"
)

// GetCryptos is a handler function that retrieves a list of cryptocurrency coins from the "Crypto" collection in the database and returns it in JSON format.
func GetCryptos(w http.ResponseWriter, r *http.Request) {

	cryptoColl := db.GetCollection("Crypto") //Uses the db.GetCollection function to obtain a collection of cryptos from the database.

	var cryptos []model.CryptoCoin //Defines an empty slice of CryptoCoin structs to store the cryptos.

	cry, err := cryptoColl.Find(context.Background(), bson.M{}) //Retrieves all documents from the collection using the Find method and stores them in a cursor variable, cry.

	if err != nil {
		return
	}

	defer cry.Close(context.Background())

	for cry.Next(context.Background()) { //terates through the cursor using a for loop and decodes each document into a CryptoCoin struct.
		var coin model.CryptoCoin
		err = cry.Decode(&coin)
		if err != nil {
			return
		}
		cryptoName := coin.Name //cryptoName gets the name of crypto being passed

		firstletter := string(cryptoName[0])           //Get the first letter of crypto name
		rest := cryptoName[1:]                         //Get the rest of the word
		newName := strings.ToLower(firstletter) + rest //Change the first letter to lowercase and add it to the rest of the word.

		price, err := price.GetPriceCoin(newName) //Put the lowercase name in the function GetPriceCoin to receive the price of the indicated crypto
		coin.Price = price                        //Updates the Price field of the struct with the price data
		if err != nil {
			return
		}
		cryptos = append(cryptos, coin) //Appends the CryptoCoin struct to the cryptos slice.
	}

	jsonByte, err := json.Marshal(cryptos) //Marshals the cryptos slice into a JSON byte slice.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") //Sets the response header.
	w.WriteHeader(http.StatusOK)                       //Set page status to OK
	w.Write(jsonByte)                                  //writes the JSON byte slice to the response body.
}

// VoteCrypto is a handler function that allows users to upvote or downvote a specific cryptocurrency coin.
func VoteCrypto(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") //Sets the Content-Type header to "application/json".

	params := r.URL.Query().Get("cryptoId") //Gets the cryptoId parameter from the query string of the request URL.

	cryptoColl := db.GetCollection("Crypto") //Uses the db.GetCollection function to obtain a collection of cryptos from the database.

	var cryptos model.CryptoCoin //Defines a CryptoCoin structs to store the cryptos.

	err := cryptoColl.FindOne(context.Background(), bson.M{"_id": params}).Decode(&cryptos) //Retrieves the corresponding CryptoCoin struct from the database based on the cryptoId.
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		fmt.Println(err)
		return
	}
	var UpVoteValue model.UpVote

	if cryptos.ID == params { //Decodes the request body into an UpVote struct to determine if the user wants to upvote (1) or downvote (0) the coin.
		_ = json.NewDecoder(r.Body).Decode(&UpVoteValue)
		fmt.Println(UpVoteValue)
		if UpVoteValue.Upvote == "down" { //If the Upvote value is 0 and if the Vote value is greater than 0, the function decrements the vote value of the coin in the database.
			if cryptos.Vote > 0 {
				err := db.DecrementVoteValue(params)
				if err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					fmt.Println(err)
					return
				}
			}
		}
		if UpVoteValue.Upvote == "up" { //If the Upvote value is 1, the function increments the vote value of the coin in the database
			err := db.UpdateVoteValue(params)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}

		fmt.Println("[upvote] Upvote value:", UpVoteValue.Upvote)
	}

	err = cryptoColl.FindOne(context.Background(), bson.M{"_id": params}).Decode(&cryptos) //Retrieves the updated CryptoCoin struct from the database based on the cryptoId
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		fmt.Println(err)
		return
	}

	cryptoName := cryptos.Name //cryptoName gets the name of crypto being passed

	firstletter := string(cryptoName[0])           //Get the first letter of crypto name
	rest := cryptoName[1:]                         //Get the rest of the word
	newName := strings.ToLower(firstletter) + rest //Change the first letter to lowercase and add it to the rest of the word.

	price, err := price.GetPriceCoin(newName) //Put the lowercase name in the function GetPriceCoin to receive the price of the indicated crypto
	cryptos.Price = price                     //Updates the Price field of the struct with the price data
	if err != nil {
		return
	}
	byte, err := json.Marshal(cryptos) //Marshals the CryptoCoin struct into JSON
	if err != nil {
		fmt.Println("[upvote] Error transforming to JSON", err)
	}

	w.WriteHeader(http.StatusOK) //Set page status to OK
	w.Write(byte)                //sends it as the response body.

}
