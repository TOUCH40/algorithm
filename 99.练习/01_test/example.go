package example

func Sum(nums ...int) int {
	sum := 0
	for num := range nums {
		sum += num
	}
	// math.
	return sum
}
