package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 1.转换之后再拼接
	var s1 = "你好吗"
	// 查看变量的内存地址
	fmt.Println(&s1)
	runeS := []rune(s1)
	runeS[2] = '啊'
	fmt.Println(string(runeS))
	//fmt.Println(&s1)

	// 2.拼接 (今天张三去上学,更改为,今天李四去上学)
	var s2 = "今天张三去上学"
	source := "张三"
	target := "李四"
	index := strings.Index(s2, source)
	sourceLen := len(source)
	start := s2[:index]
	//fmt.Printf("%T\n", start)
	//fmt.Println(start)
	end := s2[index+sourceLen:]

	var s3 bytes.Buffer
	s3.WriteString(start)
	s3.WriteString(target)
	s3.WriteString(end)
	fmt.Println(s3.String())
}
