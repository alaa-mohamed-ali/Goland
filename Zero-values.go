package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	//fmt.Printf("%v %v %v %q\n", i, f, b, s)
	//0 0 false ""

	fmt.Printf("%T %v : %T %v : %T %v : %T %q\n", i, i, f, f, b, b, s, s)
	//int 0 : float64 0 : bool false : string ""
}
