package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pablobanker/klever_test/model"
	"github.com/pablobanker/klever_test/price"
)

func Price(w http.ResponseWriter, r *http.Request) {

	//Atualiza o price a cada 5 segundos
	ticker := time.NewTicker(5 * time.Second)
	//Faz um range na api da crypto, retorna o preço da crypto solicitada.
	for range ticker.C {
		price, err := price.GetPriceCoin("bitcoin")
		if err != nil {
			fmt.Println(err)
			continue
		}
		bitcoin := model.Price{}
		bitcoin.Price = price
	}
}
