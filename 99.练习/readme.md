###  1.for-range 理解


```go
package main

import "fmt"
func main() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	fmt.Println(*m[2])// ==3,因为for...range中的val的指针分别赋给了map。且val的值最后更新为了3

}
```

