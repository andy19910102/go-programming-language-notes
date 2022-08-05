package main

import "fmt"

func main() {
	defer printTwo()
	printOne()
}

func printOne() {
	fmt.Println(1)
}

func printTwo() {
	fmt.Println(2)
}
