package main

import "fmt"

func main() {

	const pi float64 = 3.14159265

	var (
		varA = 2
		varB = 3
	)

	fmt.Println(varA, "  ", varB)

	var myName string = "Andy Huang"

	fmt.Println(len(myName))
	fmt.Println("Hello \nYour name is " + myName)

	// var isOver40 bool = true
	fmt.Printf("%.3f \n", pi) // 3.142

	fmt.Printf("%T \n", pi) // float64, return data types

	fmt.Printf("%d \n", 100) // 100
	fmt.Printf("%b \n", 100) // Get binary notation of 100
	fmt.Printf("%c \n", 44)  // Get the character code of 44
	fmt.Printf("%x \n", 17)  // Get hex code of 100
	fmt.Printf("%e \n", pi)  // Scientific notation
}
