package main

import "fmt"

func main() {
	c := []int{1, 42, 52, 11, 23, 44}
	QuickSort(c, 0, len(c)-1)
	fmt.Println(c)
}

func QuickSort(nums []int, i, j int) {
	if i >= j {
		return
	}
	p := partition(nums, i, j)
	QuickSort(nums, i, p-1)
	QuickSort(nums, p+1, j)

}

func partition(nums []int, i, j int) int {
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
