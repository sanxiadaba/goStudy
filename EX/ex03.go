package main

import (
	"fmt"
	"goStudy/utils"
)

func main() {
	var s1 = utils.UnicodeIndex("大叔", "叔")
	fmt.Println(s1)
	var s2 = utils.InsteadString("你到底", "到", "22")
	fmt.Println(s2)
}
