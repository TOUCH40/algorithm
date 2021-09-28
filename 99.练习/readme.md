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
### 2.closure 闭包
```go
type logClosure func(format string, v ...interface{})

func LoggerWrapper(logType string) logClosure {
	return func(format string, v ...interface{}) {
		fmt.Printf(fmt.Sprintf("[%s]%s", logType, format), v...)
		fmt.Println()
	}
}

func main() {
	info_logger := LoggerWrapper("INFO")
	warning_logger := LoggerWrapper("WARNING")
	info_logger("this is a %s log", "info")
	warning_logger("this is a %s log", "warning")
}
```

### 实现接口的结构体方法内部对象为指针。
### 必须以指针形式使用

```go
package main

import (
	"fmt"
)

type ex interface {
	test(int) bool
}
type example struct {
	p1 int
	p2 string
	c  func(v int) bool
	p3 struct {
		p1 int
		p2 string
	}
}

func (e example) test(c int) bool {
	e.p1 = 1
	return e.p1 == 1
}

func main() {
	var c ex = example{1, "cc",
		func(v int) bool {
			return v == 1
		},
		struct {
			p1 int
			p2 string
		}{1, "CCd"}}
	fmt.Println(c.test(1))
	fmt.Println(c.test(1))
}

```

### 反射
```go
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type MyType struct {
	i    int
	name string
}

func (mt *MyType) SetI(i int) {
	mt.i = i
}

func (mt *MyType) SetName(name string) {
	mt.name = name
}

func (mt *MyType) String() string {
	return fmt.Sprintf("%p", mt) + "--name:" + mt.name + " i:" + strconv.Itoa(mt.i)
}

func main() {
	myType := &MyType{22, "golang"}
	//fmt.Println(myType)     // 就是检查一下myType对象内容
	//println("---------------")

	// mtV := reflect.ValueOf(&myType).Elem()
	// 也可以使用
	mtV := reflect.ValueOf(myType)

	fmt.Println("Before:", mtV.MethodByName("String").Call(nil)[0])

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(18)
	mtV.MethodByName("SetI").Call(params)

	params[0] = reflect.ValueOf("reflection test")
	mtV.MethodByName("SetName").Call(params)

	fmt.Println("After:", mtV.MethodByName("String").Call(nil)[0])
}
```