package example

import (
	"testing"
)

func BenchmarkSum(t *testing.B) {
	Sum(1, 2, 3, 45, 6, 7)
	// fmt.Println(ret)
	// if ret < 10 {
	// 	t.Fatal("go to long")
	// }
}
