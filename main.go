package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pcbanker/test/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "html/index.html") }) //Creates a endpoint to /
	http.HandleFunc("/cryptos", handlers.GetCryptos)                                                               //Creates a endpoint to /cryptos
	http.HandleFunc("/upvote", handlers.VoteCrypto)                                                                //Creates a endpoint to /upvote
	fmt.Println("Opening Server")                                                                                  //Shows in the terminal that the server is working
	log.Fatal(http.ListenAndServe(":8080", nil))                                                                   //Creates a server with a portal 8585
}
