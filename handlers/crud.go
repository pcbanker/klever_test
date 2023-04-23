package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pablobanker/klever_test/model"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/pablobanker/klever_test/db"
)

func GetCryptos(w http.ResponseWriter, r *http.Request) {

	cryptoColl := db.GetCollection("Crypto")
	var cryptos []model.CryptoCoin

	cry, err := cryptoColl.Find(context.Background(), bson.M{})

	if err != nil {
		return
	}

	defer cry.Close(context.Background())

	for cry.Next(context.Background()) {
		var coin model.CryptoCoin
		err := cry.Decode(&coin)
		if err != nil {
			return
		}
		cryptos = append(cryptos, coin)
	}

	jsonByte, err := json.Marshal(cryptos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonByte)
}

func UpVoteCrypto(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query().Get("cryptoId")
	cryptoColl := db.GetCollection("Crypto")

	var cryptos model.CryptoCoin
	err := cryptoColl.FindOne(context.Background(), bson.M{"_id": params}).Decode(&cryptos)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		fmt.Println(err)
		return
	}
		var UpVoteValue model.UpVote
			if cryptos.ID == params {
				_ = json.NewDecoder(r.Body).Decode(&UpVoteValue)
				if UpVoteValue.Upvote == "0" {
					err := db.DecrementVoteValue(params)
					if err != nil {
						http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
						fmt.Println(err)
						return
					}
				}
				if UpVoteValue.Upvote == "1" {
					err := db.UpdateVoteValue(params)
					if err != nil {
						http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
						return
					}
				}
				if UpVoteValue.Upvote != "0" && UpVoteValue.Upvote != "1" {
					http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
					fmt.Println("[upvote] Error: ", err)
					return
				}
			}

	err = cryptoColl.FindOne(context.Background(), bson.M{"_id": params}).Decode(&cryptos)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	byte, err := json.Marshal(cryptos)
	if err != nil {
		fmt.Println("[upvote] Error transforming to JSON", err)
	}
	w.Write(byte)

	}

