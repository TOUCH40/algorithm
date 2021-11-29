package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "1234"
	// 方法一
	// Base64编码
	base64EncodeStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Printf("方法一,Base64编码: %v \n", base64EncodeStr)
	// Base64解码
	decodeString, err := base64.StdEncoding.DecodeString(base64EncodeStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("方法一,Base64解码: %s \n", decodeString)

	// 方法二:使用兼容URL的base64编码和解码
	toString := base64.URLEncoding.EncodeToString([]byte(str))
	fmt.Printf("方法二,Base64编码: %v \n", toString)
	bytes, _ := base64.URLEncoding.DecodeString(toString)
	fmt.Printf("方法二,Base64解码: %s \n", bytes)
}

/**输出
方法一,Base64编码: MTIzNA==
方法一,Base64解码: 1234
方法二,Base64编码: MTIzNA==
方法二,Base64解码: 1234
*/
