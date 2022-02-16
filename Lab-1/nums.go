package main

import (
	"fmt"
	"math"
)

func main() {
	var correct_int int
	var incorrect_int int
	var the_double float64

	fmt.Println("So... integers. Lets see how inputs work for this. First put in a correct int:")
	fmt.Scanln(&correct_int)
	fmt.Println("The int was: ")
	fmt.Println(correct_int)

	fmt.Println("Okay, now put in an incorrect int:")
	fmt.Scanln(&incorrect_int)
	fmt.Println("The int was: ")
	fmt.Println(incorrect_int)

	fmt.Println("One should fail... right? Nope it's 0!")
	fmt.Println(math.Sin(float64(correct_int)))

	fmt.Println("So rather than using double like in Java they use floats. Lets try")
	fmt.Scanln(&the_double)
	fmt.Println(the_double)
}
