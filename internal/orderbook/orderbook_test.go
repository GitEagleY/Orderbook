package orderbook_test

import (
	"math/rand"
	"orderbook/internal/models"
	"orderbook/internal/orderbook"
	"testing"
)

// benchmark by adding 100000 random orders and finding matches.
func BenchmarkOrderBookPerformance(b *testing.B) {

	ob := orderbook.NewOrderBook()

	// Generate and place in the order book 100000 random orders
	orders := generateRandomOrders(100000)
	for _, order := range orders {
		ob.PlaceOrder(order)
	}

	b.ResetTimer() // Reset the timer to exclude setup time from benchmark results

	// Find matches in the order book
	b.StartTimer() // Start the timer
	_, err := ob.FindMatches()
	b.StopTimer() // Stop the timer
	if err != nil {
		b.Fatalf("Error finding matches: %v", err)
	}
}

// generate a slice of n random orders for benchmarking.
func generateRandomOrders(n int) []models.Order {
	orders := make([]models.Order, n)
	for i := 0; i < n; i++ {
		// Generate random order details
		userID := rand.Int63n(1000) + 1 // between 1 and 1000
		amount := rand.Int63n(100) + 1
		price := rand.Int63n(100) + 1
		side := models.Side(rand.Intn(2))

		// Create the order
		order := models.Order{
			UserID:  userID,
			Amount:  amount,
			Price:   price,
			Side:    side,
			OrderID: int64(i + 1), // Order ID starts from 1
		}

		orders[i] = order
	}

	return orders
}
