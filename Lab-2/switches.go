package main

import (
	"fmt"
)

//NOTE: The only type of loop in go is a for loop.
func main() {
	i := 5

	fmt.Println("Let us take a number like", i)
	fmt.Println("With it we can do quite a bit with a switch statement. For example:")

	switch i {
	case 1:
		fmt.Println("This here is the number 1")
		break
	case 2:
		fmt.Println("This is the number 2")
	default:
		fmt.Println("Now now now. A number lads come on")
	}
}
