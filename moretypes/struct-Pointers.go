package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	V := Vertex{1, 2}
	p := &V
	p.X = 99

	H := Vertex{1, 2}
	c := &H
	c.Y = 101
	fmt.Println(V, H)
}
