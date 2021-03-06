package quicksort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	n := 100
	nums := GenerateS(n)
	// v := time.Now()
	QuickSort(nums, 0, n-1)
	t.Logf("%v", nums)
	// t.Logf("%d", time.Since(v).Milliseconds())
}

// 枚举 分治 递归 贪心 动态规划
func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}
