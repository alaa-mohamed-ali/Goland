package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Millisecond)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		c2 <- "TWO"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received==", msg1)
		case msg2 := <-c2:
			fmt.Println("received==", msg2)
		default:
			fmt.Println("ALAAAAA    .")
			time.Sleep(500 * time.Millisecond)
		}

	}
}
