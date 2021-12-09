package main

import (
	"fmt"
	q "main/08_quicksort"
)

func main() {
	c := q.GenerateS(100) //[]int{3, 1, 5, 6, 4, 34, 4, 2, 3, 51, 23}
	QuickSort(c, 0, len(c)-1)
	fmt.Println(c)
}

func QuickSort(nums []int, i, j int) {
	if i >= j {
		return
	}
	index := Partition2(nums, i, j)
	QuickSort(nums, i, index-1)
	QuickSort(nums, index+1, j)
}

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
