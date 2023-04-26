package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pcbanker/test/model"
)

// TestGetCryptos is a unit test function that tests the GetCryptos function, which is a handler function that retrieves a list of cryptocurrency coins.
func TestGetCryptos(t *testing.T) {

	err := godotenv.Load("./../.env") //Loads environment variables from the .env file.
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	req, err := http.NewRequest("GET", "/cryptos", nil) //Creates a new HTTP GET request to the "/cryptos" endpoint.
	if err != nil {
		log.Fatal(err)
	}
	rw := httptest.NewRecorder() //Creates a new httptest.ResponseRecorder to record the response from the handler.

	GetCryptos(rw, req) //Calls the GetCryptos handler function with the request and response recorder.

	if status := rw.Code; status != http.StatusOK { //Checks that the response status code is 200 (OK).
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var crypto []model.CryptoCoin
	err = json.Unmarshal(rw.Body.Bytes(), &crypto) //Unmarshals the JSON response into a slice of CryptoCoin structs.
	if err != nil {
		log.Fatalf("error unmarshaling JSON response: %v", err)
	}

	if len(crypto) == 0 { //Checks that the length of the slice is greater than 0.
		t.Errorf("no crypto coins returned")
	}
}

// TestUpVote is a unit test function that tests the UpVoteCrypto function, which is a handler function that increments the vote count for a specified cryptocurrency coin.
func TestUpVote(t *testing.T) {

	err := godotenv.Load("./../.env") //Loads environment variables from the .env file.
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	upvote := model.UpVote{Upvote: "up"} //Creates a new UpVote struct with a value of up.

	payload, err := json.Marshal(upvote) //Marshals the UpVote struct into a JSON payload.
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", "/upvote?cryptoId=1", bytes.NewReader(payload)) //Creates a new HTTP POST request to the "/upvote" endpoint with the cryptoId parameter set to 1 and the JSON payload in the request body.
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	rw := httptest.NewRecorder() //Creates a new ResponseRecorder to record the response from the handler

	VoteCrypto(rw, req) //Calls the UpVoteCrypto handler function with the request and response recorder.

	if status := rw.Code; status != http.StatusOK { //Checks that the response status code is 200 (OK).
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

// //TestUpVote is a unit test function that tests the UpVoteCrypto function, which is a handler function that decrements the vote count for a specified cryptocurrency coin.
func TestDownVote(t *testing.T) {

	err := godotenv.Load("./../.env") //Loads environment variables from the .env file.
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	upvote := model.UpVote{Upvote: "down"} //Creates a new UpVote struct with a value of down.

	payload, err := json.Marshal(upvote) //Marshals the UpVote struct into a JSON payload.
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", "/upvote?cryptoId=1", bytes.NewReader(payload)) ////Creates a new HTTP POST request to the "/upvote" endpoint with the cryptoId parameter set to 1 and the JSON payload in the request body.
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	rw := httptest.NewRecorder() //Creates a new ResponseRecorder to record the response from the handler

	VoteCrypto(rw, req) //Calls the UpVoteCrypto handler function with the request and response recorder.

	if status := rw.Code; status != http.StatusOK { //Checks that the response status code is 200 (OK).
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
