package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *IntHeap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}

func main() {
	h := &IntHeap{3, 4, 1, 2, 6, 7, 4}
	heap.Init(h)
	heap.Push(h, 10)
	for h.Len() > 0 {
		fmt.Printf("%d\n", heap.Pop(h))
	}
}
