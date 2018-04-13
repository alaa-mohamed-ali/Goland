package main

import "fmt"

func main() {
	/*fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}*/

	x := 42
	switch x {
	case 0:
		fmt.Printf("no\n")
	case 1:
		fmt.Printf("no\n")
	case 42:
		fmt.Printf("yes\n")
	case 43:
		fmt.Printf("no\n")
	default:
		fmt.Printf("default\n")
	}
}
