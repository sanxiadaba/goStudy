package main

import (
	"fmt"
	"strings"
)

func main() {
	// 这里没有实现unicode的位置
	str1 := "12我是珲哥"
	index1 := strings.Index(str1, "珲")
	fmt.Println(index1) //6 中文是按3个索引计算
	// 调用函数
	fmt.Println(UnicodeIndex(str1, "珲"))
}

func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}
