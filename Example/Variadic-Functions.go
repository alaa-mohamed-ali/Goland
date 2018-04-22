package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ") //[1 2 3 4]
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total) //10
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
