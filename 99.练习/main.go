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
