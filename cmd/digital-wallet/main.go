package main

import (
	"log"
	"net/http"

	"github.com/mlodovico/digital-wallet/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/wallets/", handlers.WalletHandler)
	mux.HandleFunc("/cards/", handlers.CardHandler)
	
	log.Fatal(http.ListenAndServe(":8080", mux))
}