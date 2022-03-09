package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	test(s1)
	fmt.Println(s1)
}

func test(s []int) {

	s[0] = 1000
}
