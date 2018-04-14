package main

import (
	"fmt"
)

func main() {
	names := [4]string{
		"alaa",
		"mohamed",
		"ali",
		"awad",
	}

	fmt.Println("Full Name =", names)
	//Full Name = [alaa mohamed ali awad]

	a := names[0:2]
	//[alaa mohamed]
	b := names[1:3]
	//[mohamed ali]
	fmt.Println(a, b)

	b[0] = "XXXXX"
	fmt.Println(a, b)
	//alaa XXXXX] [XXXXX ali]

	fmt.Println("Full Name =", names)
	//Full Name = [alaa XXXXX ali awad]
}
