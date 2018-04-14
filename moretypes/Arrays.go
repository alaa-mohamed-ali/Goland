package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	fmt.Println(a)

	x := [6]int{11, 22, 33, 44, 55, 66}
	fmt.Println(x)
	fmt.Println(x[3])
}
