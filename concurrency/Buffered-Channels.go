package main

import (
	"fmt"
)

func main() {
	/*c := make(chan int, 2)

	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)*/

	massages := make(chan string, 2)

	massages <- "alaa"
	massages <- "mohamed"

	fmt.Println(<-massages)
	fmt.Println(<-massages)
}
