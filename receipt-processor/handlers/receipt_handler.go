package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"receipt-processor/models"

	"github.com/google/uuid"
)

var receipts = make(map[string]models.Receipt)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	receipts[id] = receipt

	// fmt.Printf("Processed receipt ID: %s, receipt: %+v\n", id, receipt)

	response := models.ProcessResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/receipts/")
	id = strings.TrimSuffix(id, "/points")

	// fmt.Printf("Fetching points for receipt ID: %s\n", id)
	// fmt.Printf("Stored receipts: %+v\n", receipts)

	receipt, exists := receipts[id]

	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	points := calculatePoints(receipt)
	response := models.PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func calculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	alnumChars := regexp.MustCompile(`[a-zA-Z0-9]`).FindAllString(receipt.Retailer, -1)
	points += len(alnumChars)

	// Rule 2: 50 points if the total is a round dollar amount
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil && total == math.Floor(total) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil && math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Item description length multiple of 3
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the purchase date day is odd
	if date, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil && date.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time is between 2:00 PM and 4:00 PM
	if t, err := time.Parse("15:04", receipt.PurchaseTime); err == nil && t.Hour() >= 14 && t.Hour() < 16 {
		points += 10
	}

	return points
}
