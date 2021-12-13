package main

import (
	"fmt"
)

func main() {
	c := []int{42, 11, 52, 11, 23, 11, 44, 1, 44, 11}
	// QuickSort(c, 0, len(c)-1)
	BubbleSort(c)
	fmt.Println(c)
}

func BubbleSort(nums []int) {
	for j := len(nums) - 1; j > 0; j-- {
		flag := false
		for i := 0; i < j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func SelectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}

func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		// 寻找插入位置
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
}
func QuickSort(nums []int, i, j int) {
	if i >= j {
		return
	}
	p := partition2(nums, i, j)
	QuickSort(nums, i, p-1)
	QuickSort(nums, p+1, j)

}
func partition2(nums []int, i, j int) int {
	p := i + (j-i)/2
	nums[i], nums[p] = nums[p], nums[i]
	pivot := nums[i]
	for i < j {
		for i < j && nums[j] <= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] >= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	return i
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
