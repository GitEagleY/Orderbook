package orderbook

import (
	"container/heap"
	"sort"
	"sync"

	"github.com/GitEagleY/Orderbook/internal/models"
)

type OrderBook struct {
	BuyOrders  OrderHeap // Priority queue for buy orders
	SellOrders OrderHeap // Priority queue for sell orders
	OrderID    int64     // Counter for generating unique order IDs
	mutex      sync.Mutex
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		BuyOrders:  make(OrderHeap, 0),
		SellOrders: make(OrderHeap, 0),
		OrderID:    0,
	}
}

func (ob *OrderBook) PlaceOrder(order models.Order) ([]models.BalanceChange, error) {
	ob.mutex.Lock()
	defer ob.mutex.Unlock()

	ob.OrderID++
	order.OrderID = ob.OrderID

	if order.Side == models.Buy {
		heap.Push(&ob.BuyOrders, order)
	} else {
		heap.Push(&ob.SellOrders, order)
	}
	// Calculate Total Price
	totalPrice := order.Amount * order.Price

	balanceChanges := make([]models.BalanceChange, 0)
	if order.Side == models.Buy {
		balanceChanges = append(balanceChanges, models.BalanceChange{UserID: order.UserID, Value: -totalPrice, Currency: "USD"})
		balanceChanges = append(balanceChanges, models.BalanceChange{UserID: order.UserID, Value: order.Amount, Currency: "UAH"})
	} else {
		balanceChanges = append(balanceChanges, models.BalanceChange{UserID: order.UserID, Value: totalPrice, Currency: "USD"})
		balanceChanges = append(balanceChanges, models.BalanceChange{UserID: order.UserID, Value: -order.Amount, Currency: "UAH"})
	}

	return balanceChanges, nil
}

func (ob *OrderBook) FindMatches() ([]models.Order, error) {
	var matches []models.Order

	ob.mutex.Lock()
	defer ob.mutex.Unlock()

	// Sort buy and sell orders by price
	sort.SliceStable(ob.BuyOrders, func(i, j int) bool {
		return ob.BuyOrders[i].Price < ob.BuyOrders[j].Price
	})
	sort.SliceStable(ob.SellOrders, func(i, j int) bool {
		return ob.SellOrders[i].Price < ob.SellOrders[j].Price
	})

	if len(ob.BuyOrders) == 0 || len(ob.SellOrders) == 0 {
		return matches, nil
	}

	// Perform binary search for matching orders
	for _, buyOrder := range ob.BuyOrders {
		left := 0
		right := len(ob.SellOrders) - 1

		for left <= right {
			mid := left + (right-left)/2

			if ob.SellOrders[mid].Price <= buyOrder.Price && ob.SellOrders[mid].Amount == buyOrder.Amount {
				matches = append(matches, buyOrder, ob.SellOrders[mid])
				break
			} else if ob.SellOrders[mid].Price < buyOrder.Price {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return matches, nil
}

func (ob *OrderBook) GetAllOrders() []models.Order {
	ob.mutex.Lock()
	defer ob.mutex.Unlock()

	allOrders := make([]models.Order, 0)

	allOrders = append(allOrders, ob.SellOrders...)
	allOrders = append(allOrders, ob.BuyOrders...)

	return allOrders
}
