package main

//func faa(arg interface{}) {
//	fmt.Println(arg)
//	fmt.Printf("%T\n", arg)
//}

func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	strSlice := []string{"a", "b", "c", "d"}
	boolSlice := []bool{true, true, false, true}
	faa(intSlice)
	faa(strSlice)
	faa(boolSlice) // 能够成功输出，不会报错
}
