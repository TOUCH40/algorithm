package quicksort

import (
	"math/rand"
	"time"
)

func GenerateS(n int) (ret []int) {
	ret = make([]int, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		ret = append(ret, rand.Intn(100))
	}
	return
}

func QuickSort(nums []int, i, j int) {
	if i >= j {
		return
	}
	index := Partition(nums, i, j)
	QuickSort(nums, i, index-1)
	QuickSort(nums, index+1, j)
}

// 填坑
func Partition(nums []int, i, j int) int {
	pivot := nums[i]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	return i
}

func Partition2(nums []int, i, j int) int {
	pivot := nums[i]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[j] = pivot
	return j
}
