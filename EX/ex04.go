package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := s[1:3]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}
