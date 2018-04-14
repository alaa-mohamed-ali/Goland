package main

import (
	"fmt"
)

func main() {
	x := [6]int{2, 3, 5, 7, 11, 13}
	//var s []int = x[1:4]
	var s []int = x[2:4]
	fmt.Println(s)
}
