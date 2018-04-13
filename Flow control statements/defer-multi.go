package main

import (
	"fmt"
)

func main() {

	//ONE
	fmt.Println("counting")

	for i := 1; i < 10; i++ {
		//THREE
		defer fmt.Println(i)
	}

	//TOW
	fmt.Println("done")

	/*
		counting
		done
		9
		8
		7
		6
		5
		4
		3
		2
		1

	*/
}
