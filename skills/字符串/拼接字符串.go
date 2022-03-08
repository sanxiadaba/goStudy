package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var s1 string = "我是字符串1"
	var s2 string = "我是字符串2"

	// 1.使用+号来拼接 (较低效)
	fmt.Println(s1 + s2)

	// 2.使用专门实现拼接的类来拼接
	var s3 bytes.Buffer
	s3.WriteString(s1)
	s3.WriteString(s2)
	fmt.Printf("%T,%s\n", s3, s3)
	// 转换为字符串
	s4 := s3.String()
	fmt.Printf("%T,%s\n", s4, s4)

	// 3.使用join拼接
	hello := "hello"
	world := "world"
	newHello := strings.Join([]string{hello, world}, ",")
	fmt.Println(newHello)
}
