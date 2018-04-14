package main

import (
	"fmt"
)

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	c := []bool{true, false, true, true, false, true}
	fmt.Println(c)

	s := []struct {
		X int
		Y bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
}
