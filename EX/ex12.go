package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4}
	s2 := s1[2:3]
	var s3 = make([]int, 1, 10)
	copy(s3, s1[2:3])
	s2[0] = 100
	fmt.Println(s1, s2, s3)
}
