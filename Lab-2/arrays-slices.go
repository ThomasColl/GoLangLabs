package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr [10]int

	fmt.Println("When an array is first defined it is empty", arr)

	arr[6] = 69
	fmt.Println("But as you see once you add in a value it is no longer all empty", arr[6])
	fmt.Println("You can also show off a single entry", arr[5], arr[6])

	var two_d_array [3][3]int

	for i := 0; i < len(two_d_array); i++ {
		for j := 0; j < len(two_d_array[i]); j++ {
			two_d_array[i][j] = rand.Intn(2)
		}
	}

	fmt.Println(two_d_array)

	fmt.Println("-----------------------------------------")
	fmt.Println("Now for Slices!")
	fmt.Println("-----------------------------------------")

	slice := make([]int, 5)
	fmt.Println("Here is an empty slice:", slice)

	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(10)
	}

	fmt.Println("Now here is it when filled:", slice)

	slice = append(slice, 4, 2, 0)

	fmt.Println("Appended you can add 420", slice)
}
