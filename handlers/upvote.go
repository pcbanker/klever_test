package handlers

import (
	"fmt"
	"net/http"

	"github.com/pablobanker/klever_test/db"

	"github.com/google/uuid"
)

func UpVoteCoin(w http.ResponseWriter, r *http.Request) {

	//Cria um variavel para criar um id aleatorio para o user
	id := uuid.New()
	//Ele adiciona o id no banco de dados
	err := db.CreateUser(id.String())
	if err != nil {
		fmt.Println("Error inserting the user id in the DB", err.Error())
	}
	fmt.Println("User Added")

	//Id recebido do botao clicado pelo user

	var crypto = "4"
	//Procura no database se a crypto existe
	_, err = db.FindByCryptoId(crypto)
	if err != nil {
		//Caso não exista ele me retorna um error
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Println("[upvote] Error finding the crypto", err.Error())
		return
	}
	// Procura a Id da crypto no database do user para ver se ele votou
	_, err = db.FindVoteCrypto(id.String(), crypto)
	//Caso dê error significa que ele não votou ainda nessa crypto
	if err != nil {
		//Adiciona o voto no database da Crypto
		err = db.UpdateVoteValue(crypto)
		if err != nil {
			return
		}
		//Adiciona o voto no database do User
		_, err = db.UpdateVoteValueUser(id.String(), crypto)
		if err != nil {
			return
		}
	}
	//Caso não der error significa que ele votou naquela crypto
	//Então irá remover 1 voto da crypto clicada
	err = db.DecrementVoteValue(crypto)
	if err != nil {
		return
	}
	//Irá remover o voto e a crypto do database do User
	_, err = db.DecrementVoteValueUser(id.String(), crypto)
	if err != nil {
		return
	}
}
