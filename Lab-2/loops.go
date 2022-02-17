package main

import (
	"fmt"
	"strconv"
)

//NOTE: The only type of loop in go is a for loop.
func main() {
	i := 0

	for i < 3 {
		fmt.Println("i is : " + strconv.Itoa(i))
		i = i + 1
	}

	sum := 0
	for i = 0; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	for {
		fmt.Println("Lets take a break")
		break
	}

	for i = 10; i >= 1; i-- {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
