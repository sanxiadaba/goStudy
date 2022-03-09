package main

import "fmt"

func main() {
	// 切片的初始化
	s1 := []int{1, 2, 3}
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Printf("%T\n", s1)

	// 初始化时定义长度与容量
	s3 := make([]int, 1, 2)
	fmt.Println(s3, len(s3), cap(s3))

	// 数组的切片
	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("原数组：", a1)
	fmt.Println("对数组进行截取：")
	//如果指定max，max的值最大不能超过截取对象（数组、切片）的容量
	s2 := a1[2:5:9] //max:9  low：2  high;5  len:5-2(len=high-low)  cap:9-2(cap=max-low)
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2))

	// 注意,这样初始化的slice不为nil
	var s4 []string
	// 这种情况下是true
	fmt.Println(s4 == nil)
	var s5 = []string{}
	// 这种情况下是false
	fmt.Println(s5 == nil)

	// 切片的引用传递
	var a2 = [5]int{1, 2, 3, 4, 5}
	// 注意这里s6的容量
	var s6 = a2[1:3]
	s6 = s6[0:cap(s6)]
	fmt.Println(s6) // [2 3 4 5]

	var s7 = a2[1:3]
	// 看看数组,切片里的元素是不是指向同一个内存地址
	// 事实证明指向的是相同的地址
	fmt.Printf("%p\n", &(a2[1]))
	fmt.Printf("%p\n", &(s7[0]))
	fmt.Println("改变前的a2的值: ", a2)
	s7[0] = 100
	fmt.Println("改变后的a2的值: ", a2)

	fmt.Println("==================================")
	// go语言切片的复制 // 注意这里地址已经改变了
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	slice2[0] = 10000
	fmt.Println(slice1, slice2)
	fmt.Printf("%p,%p\n", &slice1[0], &slice2[0])
}
