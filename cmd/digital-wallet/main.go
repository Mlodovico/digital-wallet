package main

import (
	"log"
	"net/http"

	"github.com/mlodovico/digital-wallet/internal/handlers"
)

func main() {
    http.HandleFunc("/digital-walllet", handlers.WalletHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}