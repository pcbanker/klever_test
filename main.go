package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pablobanker/klever_test/handlers"
)

func main() {

	http.HandleFunc("/cryptos", handlers.GetCryptos)
	http.HandleFunc("/upvote", handlers.UpVoteCrypto)
	fmt.Println("Opening Server")
	log.Fatal(http.ListenAndServe(":8585", nil))
}
