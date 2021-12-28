package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("aaababdefaaababdefg", "ababdefg"))
}

func strStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	// 构建 next 数组，数组长度为匹配串的长度（next 数组是和匹配串相关的）
	next := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[j] != needle[i] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && needle[j] != haystack[i] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - j + 1
		}
	}
	return -1
}
