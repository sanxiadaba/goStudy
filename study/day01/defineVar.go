package main

import "fmt"

// 定义变量

func main() {
	// 1.定义后再赋值
	var id int
	id = 1
	// 2.定义时赋值
	var age int = 10
	// 3.不定义类型,直接赋值
	var name string = "名字"
	// 4.批量声明变量
	var (
		password string  // 输出为空
		filename         = "文件名"
		headline         = "标题"
		money    float32 = 1.0
	)

	// 5.简短格式定义 必须定义在函数内部
	zone := 10

	// 6.可以重复定义的情况(左边只要有一个是新变量即可)
	a, b := 1, 2
	b, c := 2, 3

	// 7.交换变量
	var d = 10
	var e = 20
	d, e = e, d

	// 输出变量
	fmt.Println(id)
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(password)
	fmt.Println(headline)
	fmt.Println(filename)
	fmt.Println(zone)
	fmt.Println(a, b, c)
	fmt.Println(d, e)

	// 输出变量的类型
	fmt.Printf("%T\n", id)

	// 格式化输出
	fmt.Printf("%s,%d,%f\n", headline, age, money)

}
