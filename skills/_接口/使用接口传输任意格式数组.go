package main

import (
	"fmt"
	"reflect"
)

func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	strSlice := []string{"a", "b", "c", "d"}
	boolSlice := []bool{true, true, false, true}

	faa(intSlice)
	faa(strSlice)
	faa(boolSlice)
}

/*能够正常输出,但是这种情况不能使用for循环对数组进行遍历*/
//func faa_1(arg interface{}) {
//	fmt.Println(arg, "\t", reflect.TypeOf(arg))
//	for _,v := range arg {
//		fmt.Println(v)
//	}
//}

func CreateAnyTypeSlice(slice interface{}) ([]interface{}, bool) {
	val, ok := isSlice(slice)

	if !ok {
		return nil, false
	}

	sliceLen := val.Len()

	out := make([]interface{}, sliceLen)

	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}

	return out, true
}

// 利用反射判断是否为slcie数据
func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == reflect.Slice {
		ok = true
	}
	return val, ok
}

func faa(arg interface{}) {
	slice, ok := CreateAnyTypeSlice(arg)
	if !ok {
		return
	}
	for index, value := range slice {
		fmt.Println(index, value)
	}
}
