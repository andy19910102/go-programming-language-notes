package main

import "fmt"

func main() {
	x := 0
	changeXvalNow(&x)
	fmt.Println("x =", x)
	fmt.Println("Memory Address for x= =", &x)

	yPtr := new(int)
	changeYValNow(yPtr)
	fmt.Println("y =", *yPtr)
	fmt.Println("Memory Address for yPtr=", yPtr)
}

func changeXvalNow(x *int) {
	*x = 2
}

func changeYValNow(yPtr *int) {
	*yPtr = 100
}
