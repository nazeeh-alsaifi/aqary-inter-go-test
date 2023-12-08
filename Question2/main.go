package main

import (
	"container/heap"
	"fmt"
)

type CharFrequency struct {
	char      byte
	frequency int
}

type MaxHeap []CharFrequency

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].frequency > h[j].frequency }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(CharFrequency))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func rearrangeString(s string) string {
	charCount := make(map[byte]int)

	// Count the frequency of each character
	for i := 0; i < len(s); i++ {
		charCount[s[i]]++
	}

	// Create a max heap and push character frequencies
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for char, count := range charCount {
		heap.Push(maxHeap, CharFrequency{char, count})
	}

	result := make([]byte, 0, len(s))

	// Build the result string by alternately appending characters
	for maxHeap.Len() >= 2 {
		// Pop two characters with the highest frequency
		cf1 := heap.Pop(maxHeap).(CharFrequency)
		cf2 := heap.Pop(maxHeap).(CharFrequency)

		// Append the characters to the result string
		result = append(result, cf1.char, cf2.char)

		// Decrement the frequency and push back to the heap if the count is not zero
		if cf1.frequency-1 > 0 {
			heap.Push(maxHeap, CharFrequency{cf1.char, cf1.frequency - 1})
		}
		if cf2.frequency-1 > 0 {
			heap.Push(maxHeap, CharFrequency{cf2.char, cf2.frequency - 1})
		}
	}

	// If there is a single character with remaining frequency, append it to the result
	if maxHeap.Len() > 0 {
		cf := heap.Pop(maxHeap).(CharFrequency)
		result = append(result, cf.char)
	}

	// Check if the rearrangement is possible
	if len(result) != len(s) {
		return ""
	}

	return string(result)
}

func main() {
	fmt.Println(rearrangeString("aab"))   // Output: "aba"
	fmt.Println(rearrangeString("aaab"))  // Output: ""
	fmt.Println(rearrangeString("aaabb")) // Output: "ababa"

}
