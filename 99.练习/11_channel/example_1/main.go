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
	
}

func kthLargest(root *TreeNode, k int) int {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
	cancel()
	}()
	
	ch := make(chan int)
	
	go work(ctx, root, ch)
	for i := 0; i < k-1; i++ {
	<-ch
	}
	
	return <-ch
	}
	
	func work(ctx context.Context, root *TreeNode, ch chan<- int) {
	if root == nil {
	return
	}
	work(ctx, root.Right, ch)
	select {
	case ch <- root.Val:
	case <-ctx.Done():
	return
	}
	work(ctx, root.Left, ch)
	}
