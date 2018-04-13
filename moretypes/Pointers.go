package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println("p == ", *p)
	//p == 42

	*p = 21
	fmt.Println("i == ", i)
	//i == 21

	x := &j
	fmt.Println("x == ", *x)
	//x == 2701

	*x = *x / 37
	fmt.Println("j == ", j)
	//j ==  73

	alaa := 55
	c := &alaa
	fmt.Println("c == ", *c)
	*c = 33
	fmt.Println("alaa == ", alaa)

	/**p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37*/
	//	fmt.Println(j)
}
