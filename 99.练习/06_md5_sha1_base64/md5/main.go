package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	str := "Hello Word"
	// 方法一
	sum := md5.Sum([]byte(str))
	// 转成16进制
	md51 := fmt.Sprintf("%x", sum)
	fmt.Println("方法一: " + md51)

	// 方法二
	hash := md5.New()
	// 将str写入到hash
	_, err := io.WriteString(hash, str)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 转成16进制
	md52 := fmt.Sprintf("%x", hash.Sum(nil))
	fmt.Println("方法二: " + md52)
}

/** 输出
方法一: ed0a96e83ab7b0910fcbcc131b2e6b82
方法二: ed0a96e83ab7b0910fcbcc131b2e6b82
*/
