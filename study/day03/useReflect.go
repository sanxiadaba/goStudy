package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1.查看数据的类型
	a := 1
	b := []int{1, 2, 3}
	type cat struct {
	}
	c := cat{}
	// int []int
	// 查看数据的类型
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(c).Name(), reflect.TypeOf(c).Kind())
	fmt.Println(reflect.TypeOf(b).Kind())

}
