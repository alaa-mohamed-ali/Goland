package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	//fmt.Printf("Type : %T Value : %v\n", f, f)
	//Type : float64 Value : 5

	fmt.Println(x, y, z)
	// 3 4 5
}
