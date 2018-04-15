package main

import "fmt"

type Vertex struct {
	//Lat, Long float64
	Lat, Long int
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40, -74,
	}
	fmt.Println(m["Bell Labs"])
}
