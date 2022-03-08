package main

import "fmt"

func main() {
	str := "a中文cd"
	str = string([]rune(str)[:3])
	fmt.Println(str)
}
