package main

import (
	"fmt"
)

func main() {
	//var a []int
	a := make([]int, 0, 20)
	a = append(a, 1) // 追加1个元素
	fmt.Println(a)
	a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
	fmt.Println(a)
	// 注意: 切片拼接切片只能有两个参数,且第二个参数后加...
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	fmt.Println(a)
}
