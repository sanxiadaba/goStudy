package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int 转string
	var str1 string
	a1 := 10
	str1 = strconv.FormatInt(int64(a1), 10)
	fmt.Printf("%T\n", str1)

	// float 转string
	// 第一个参数要是float类型 当第二个参数为f时第三个参数表示小数点后保留的位数,为g时总共的位数
	println(strconv.FormatFloat(3.1415926, 'f', 5, 64)) // 3.14159

	// 字符转整形
	a, _ := strconv.Atoi("001")
	println(a)

	// 字符串转整形
	float, _ := strconv.ParseFloat("123.213", 64)
	fmt.Printf("%T", float)
	println(float + 123.869)
	var lin float32
	lin = 1000.0
	println(lin)
}
