package main

import "fmt"

func main() {
	str := "Golang，你好吗？"

	//以Unicode字符方式遍历
	for i, ch := range str {
		// 1.使用%c格式输出
		fmt.Printf("%c\n", ch)
		// 2.转换为string
		fmt.Printf("%.2d：%s\n", i, string(ch)) //ch为代表Unicode字符的rune类型
	}
}
