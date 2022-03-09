package main

import "fmt"

func main() {
	//var a1 = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//var s1 = a1[3:5:5]
	//fmt.Println(len(s1), cap(s1))

	var a1 = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s1 = a1[3:5]
	fmt.Println(len(s1), cap(s1))
	s1 = s1[0:cap(s1)]
	fmt.Println(s1)
}
