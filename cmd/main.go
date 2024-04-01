package main

import (
	"fmt"
	"log"
	"orderbook/internal/handlers"
	"orderbook/internal/orderbook"
	"os"
)

func main() {
	ob := orderbook.NewOrderBook()
	for {
		fmt.Println("1. Place Order")
		fmt.Println("2. See Matching Orders")
		fmt.Println("3. Display All Orders")
		fmt.Println("4. Exit")
		fmt.Print("Select an option: ")

		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			log.Println("Error reading input:", err)
			continue
		}

		switch choice {
		case 1:
			handlers.PlaceOrder(ob)
		case 2:
			handlers.DisplayMatchingOrders(ob)
		case 3:
			handlers.DisplayAllOrders(ob)
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
