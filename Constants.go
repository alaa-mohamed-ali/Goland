package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "alaa mohamed ali"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	/*
		Hello alaa mohamed ali
		Happy 3.14 Day
		Go rules? true
	*/

	fmt.Println(Pi, World, Truth)
	//3.14 alaa mohamed ali true
}
