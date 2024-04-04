# Orderbook

## Prerequisites

Before running the program, ensure you have Go installed on your system. You can download and install Go from the [official website](https://go.dev/dl/).

## How to Run

1.  Clone this repository to your local machine.
2.  Navigate to the project directory in your terminal.
3.  Run the following command to build and execute the program:

Linux/Mac:

    go run cmd/main.go

Win 10:

    go run cmd\main.go

## Data Structures

- **OrderBook**: Implemented using two priority queues to store buy and sell orders separately. This allows efficient retrieval of the highest priority orders (with lowest price).
- **OrderHeap**: Used to represent the priority queue (heap) for orders. It is a slice of `models.Order` structs that implements the `heap.Interface`.

## Algorithm for Matching Orders

1. Locking the Mutex: The function begins by acquiring a lock on the mutex (ob.mutex.Lock()) to ensure exclusive access to the OrderBook data structure. This prevents concurrent modifications to the data.

2. Sorting: The buy and sell orders are sorted by price using sort.SliceStable. Sorting is crucial for binary search to work efficiently because binary search requires the data to be in a sorted order.

3. Binary Search: The function iterates over each buy order in ob.BuyOrders and performs a binary search for matching sell orders in ob.SellOrders. Here's how binary search is applied:
    1. The function initializes two pointers, left and right, to define the search interval within the ob.SellOrders array.
    2. Inside a nested loop, the function calculates the middle index (mid) of the search interval and compares the price of the sell order at mid with the price of the current buy order.
    3. Depending on the comparison, the search interval is narrowed down by updating the left or right pointer.
    4. If a matching sell order is found (i.e., price and amount match), both the buy and sell orders are appended to the matches slice.
    5. The binary search continues until either a match is found or the search interval is empty.

4. Returning Results: After all buy orders are processed, the function returns the matches slice containing pairs of matching buy and sell orders. It also releases the lock on the mutex (ob.mutex.Unlock()).

## Concurrency

Concurrency is implemented using mutexes to ensure thread safety when accessing shared data structures, such as the order book.

## Why These Choices?

- **Priority queues**: Priority queues are well-suited for order book implementations as they efficiently handle inserting new orders and retrieving the highest priority orders.
- **Matching Algorithm**: The chosen algorithm considers both price and amount, ensuring that only orders where the seller can fulfill the buyer's request are considered as matches.
- **Goroutines and maps**: Also while i was trying to implement maps and goroutines benchmark time was increasing dramatically so i decided to not use them for this project

## Benchmarking

The project includes benchmark tests to measure the performance of critical components, such as finding matches.
To run the benchmark tests, execute the following command:

Linux/Mac:

    cd internal/orderbook && go test -bench=.

Win10:

    cd internal\orderbook && go test -bench=.

personally I receive 0,8s of time by sorting 100000 matches

## Efficiency and Big O Notation

PlaceOrder:  
- Insertion into the heap: O(log n)
- Calculating balance changes: O(1)
FindMatches:
- Nested loop iterating over all buy and sell orders with binary search: O(log n)

GetAllOrders:
- Iterating over all buy and sell orders: O(n)

## Time

~8-9 hours spent solving the problem
