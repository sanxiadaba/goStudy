package main

import "fmt"

func main() {
	myfunc(1, 2, 3)
}

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
