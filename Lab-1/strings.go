package main

import "fmt"

func main() {
	var first string
	var second string
	const a_constant_string = "Constants, constants never change"

	fmt.Println("I'm gonna ask you to type a few things in okay? \n")

	fmt.Println("So what is the first thing you wanna say?")
	fmt.Scanln(&first)
	fmt.Println("The first thing you typed is: " + first)

	fmt.Print("Interesting........ but what about a second? \n")
	fmt.Scanln(&second)
	fmt.Println("Oh so the second is " + second)

	fmt.Println("I would like to teach you about constants now but..... " + a_constant_string)
}
