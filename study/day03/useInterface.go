package main

import "fmt"

func main() {
	c := cat{}
	d := dog{}
	da(c)
	da(d)

	// 注意,任何接口都实现了空接口,所以可以传递任意类型
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	strSlice := []string{"a", "b", "c", "d"}
	boolSlice := []bool{true, true, false, true}
	faa(intSlice)
	faa(strSlice)
	faa(boolSlice) // 能够成功输出，不会报错
}

// 定义一个接口
// 注意: 要实现接口里的所有方法才算是这个接口
type sayer interface {
	say()
}

type cat struct {
}

type dog struct {
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func da(arg sayer) {
	arg.say()
}

func faa(arg interface{}) {
	fmt.Println(arg)
	fmt.Printf("%T\n", arg)
}
