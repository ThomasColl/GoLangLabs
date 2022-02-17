package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Println("Enter a non zero number to see if it is even or odd:")
	fmt.Scanln(&num)
	if num == 0 {
		fmt.Println("Now now now. A non zero number please!")
	} else if num%2 == 0 {
		fmt.Println("Ah " + strconv.Itoa(num) + " is even")
	} else {
		fmt.Println("Ah " + strconv.Itoa(num) + " is odd")
	}
}
