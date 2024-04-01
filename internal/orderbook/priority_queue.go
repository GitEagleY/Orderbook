package orderbook

import (
	"github.com/GitEagleY/Orderbook/internal/models"
)

// OrderHeap is a slice of Order structs that implements the heap.Interface.
// It represents a priority queue where the highest priority item is at the root.
type OrderHeap []models.Order

func (heap OrderHeap) Len() int { return len(heap) }

// reports whether the element with index i should sort before the element with index j.
func (heap OrderHeap) Less(i, j int) bool {
	return heap[i].Price < heap[j].Price
}

// swaps elements
func (heap OrderHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

// adds an item to the heap.
func (heap *OrderHeap) Push(x interface{}) {
	// Convert the item to Order + append it to the slice.
	*heap = append(*heap, x.(models.Order))
}

// removes the highest priority item from the heap and returns it.
func (heap *OrderHeap) Pop() interface{} {
	old := *heap
	lngth := len(old)
	last := old[lngth-1]     // Get the last item
	*heap = old[0 : lngth-1] // Remove the last item from the slice
	return last              // Return the removed item
}
