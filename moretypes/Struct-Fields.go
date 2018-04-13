package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	V := vertex{1, 2}
	V.X = 4
	V.Y = 9
	fmt.Println(V.X, V.Y)
	//fmt.Println(vertex{1, 2})
}
