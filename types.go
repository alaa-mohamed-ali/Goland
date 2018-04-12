package main

import (
	"fmt"
	"math/cmplx"
)

var (
	i bool       = false
	j uint64     = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type : %T Value : %v\n", i, i)
	fmt.Printf("Type %T Value %v\n", j, j)
	fmt.Printf("Type %T Value %v\n", z, z)

	/*
		Type : bool Value : false
		Type uint64 Value 18446744073709551615
		Type complex128 Value (2+3i)
	*/
}
