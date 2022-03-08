package main

import "fmt"

func main() {
	var a int = 10 /* 声明实际变量 */
	// 以指针类型输出
	fmt.Printf("变量的地址: %p\n", &a)

	// 指针的基本使用
	var ip *int /* 声明指针变量 */
	ip = &a     /* 指针变量的存储地址 */

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	// 使用指针修改值
	s := 10
	fmt.Printf("修改前: %d", s)
	modifyUsePoint(&s)
	fmt.Printf("修改后: %d\n", s)

	// 使用new创建指针
	var s1 = new(string)
	*s1 = "指针赋值"
	fmt.Println(s1)  //这里输出的是指针的地址
	fmt.Println(*s1) //这里输出的是指针的值
}

func modifyUsePoint(ptr *int) {
	*ptr = 100
}
