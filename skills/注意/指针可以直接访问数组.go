package main

import "fmt"

func main() {
	// 指针直接访问数组
	var arrPtr *[4]int           // 创建一个指针 arrPtr，指向一个数组
	var arr = [4]int{1, 2, 3, 4} // 创建一个数组并初始化
	arrPtr = &arr                // 将数组 arr的地址赋值给arrPtr
	fmt.Println("将 arr 的内存地址赋值给数组指针 arrPtr,   arrPtr=", arrPtr)
	fmt.Println(arrPtr[0])
}
