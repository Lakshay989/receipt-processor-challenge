package main

import (
	"log"
	"net/http"

	"receipt-processor/handlers"
)

func main() {
	http.HandleFunc("/receipts/process", handlers.ProcessReceipt)
	http.HandleFunc("/receipts/", handlers.GetPoints)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
