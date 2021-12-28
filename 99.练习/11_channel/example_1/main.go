package main

import "context"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
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
