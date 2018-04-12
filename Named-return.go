package main

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	// 17 * 4 / 9 = 7
	y = sum - x
	// 17 - 7 = 10
	return
}

func main() {
	fmt.Println(split(17))
}
