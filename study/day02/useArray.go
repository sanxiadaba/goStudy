package main

import "fmt"

func main() {
	// 1.定义一定长度空数组
	var a1 [5]float32
	// 不赋值的话默认为0
	println(a1[0]) // +0.000000e+000

	//  2.定义数组时赋值
	var a2 = [5]int{1, 2, 3, 4, 5}
	// println(a2) // 数组不能打印(不能只使用println  可以用fmt.println打印)
	println(a2[0], len(a2))

	// 3.定义时自动推断数组长度
	var a3 = [...]int{2, 4, 6}
	//println(a3)
	// 使用格式化输出来输出数组
	fmt.Printf("%v\n", a3)

	// 4.给指定长度的数组起别名,方便定义
	type arr3 [3]int
	var a4 arr3 = arr3{1, 4, 5}
	println(a4[0])

	// 从数组里取值
	for i, i2 := range a2 {
		println(i, i2)
	}

	// 定义多维数组
	var a5 = [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	fmt.Printf("%v", a5)

	// 多维数组的使用

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// Step 3: 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Step 4: 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])
	fmt.Println("===========================")

	// 注意,数组是值传递
	//常见两种初始化方式
	//var b = [...]int{1, 2, 3}
	var b = [3]int{1, 2, 3}
	updateArr(b)
	fmt.Println(b)
	updateArrPoint(&b)
	fmt.Println(b)
	//计算数组长度和容量
	fmt.Println(len(b))
	fmt.Println(cap(b))

	fmt.Println("===============")
	// 数组的切片
	a6 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("原数组：", a6)

	fmt.Println("对数组进行截取：")
	//如果指定max，max的值最大不能超过截取对象（数组、切片）的容量
	s7 := a6[2:5:9] //max:9  low：2  high;5  len:5-2(len=high-low)  cap:9-2(cap=max-low)
	fmt.Println(s7)
	fmt.Println(len(s7), cap(s7))
}

//值传递,传的是副本
func updateArr(b [3]int) {
	b[0] = 3
}

//传指针,[3]int是一个类型
func updateArrPoint(b *[3]int) {
	fmt.Printf("%T\n", b)
	// 标准写法应该这样
	(*b)[0] = 3
	// 在go中,可以这样
	b[0] = 3
}
