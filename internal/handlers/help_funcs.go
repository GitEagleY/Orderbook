package handlers

import (
	"fmt"
	"log"
	"orderbook/internal/models"
	"orderbook/internal/orderbook"
)

func PlaceOrder(ob *orderbook.OrderBook) {
	var userID, amount, price int64
	var side models.Side

	fmt.Print("Enter User ID: ")
	_, err := fmt.Scanf("%d", &userID)
	if err != nil {
		log.Println("Error reading User ID:", err)
		return
	}

	fmt.Print("Enter Amount: ")
	_, err = fmt.Scanf("%d", &amount)
	if err != nil {
		log.Println("Error reading Amount:", err)
		return
	}

	fmt.Print("Enter Price: ")
	_, err = fmt.Scanf("%d", &price)
	if err != nil {
		log.Println("Error reading Price:", err)
		return
	}

	fmt.Print("Enter Side (1 for buy or 0 for sell)(default: buy): ")
	_, err = fmt.Scanf("%d", &side)
	if err != nil {
		log.Println("Error reading Side:", err)
		return
	}

	order := models.Order{
		UserID: userID,
		Amount: amount,
		Price:  price,
		Side:   side,
	}

	balanceChanges, err := ob.PlaceOrder(order)
	if err != nil {
		fmt.Println("Error placing order:", err)
		return
	}

	fmt.Println("Order placed successfully!")
	fmt.Println("Expected Balance Changes:")
	for _, bc := range balanceChanges {
		fmt.Printf("UserID: %d, Value: %d %s\n", bc.UserID, bc.Value, bc.Currency)
	}
}

func DisplayMatchingOrders(ob *orderbook.OrderBook) {
	matches, err := ob.FindMatches()
	if err != nil {
		fmt.Println("Error finding matches:", err)
		return
	}

	if len(matches) == 0 {
		fmt.Println("No matching orders found.")
		return
	}

	fmt.Println("Matching Orders:")
	for i := 0; i < len(matches); i += 2 {
		buyOrder := matches[i]
		sellOrder := matches[i+1]
		fmt.Println("-------------------------")
		fmt.Printf("Buy Order: OrderID: %d, UserID: %d, Amount: %d, Price: %d, Side: %s\n",
			buyOrder.OrderID, buyOrder.UserID, buyOrder.Amount, buyOrder.Price, buyOrder.Side)
		fmt.Printf("Sell Order: OrderID: %d, UserID: %d, Amount: %d, Price: %d, Side: %s\n",
			sellOrder.OrderID, sellOrder.UserID, sellOrder.Amount, sellOrder.Price, sellOrder.Side)
		fmt.Println("-------------------------")
	}
}

func DisplayAllOrders(ob *orderbook.OrderBook) {
	allOrders := ob.GetAllOrders()

	if len(allOrders) == 0 {
		fmt.Println("No orders found.")
		return
	}

	fmt.Println("All Orders:")
	for _, order := range allOrders {
		fmt.Println("-------------------------")
		fmt.Printf("OrderID: %d, UserID: %d, Amount: %d, Price: %d, Side: %s\n",
			order.OrderID, order.UserID, order.Amount, order.Price, order.Side)
		fmt.Println("-------------------------")
	}
}
