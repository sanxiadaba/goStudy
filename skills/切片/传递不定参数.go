package main

import "fmt"

func main() {
	myfunc(1, 2, 3)
}

func myfunc(s int, args ...int) {
	// 此时传进来的args是一个切片
	fmt.Printf("%T\n", args)
	fmt.Println(len(args), cap(args))
	for i, arg := range args {
		fmt.Println(i, arg)
	}
	fmt.Println("s是", s)
}
