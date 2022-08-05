package main

import "fmt"

func main() {
	demPanic()
}

func demPanic() {

	defer func() {
		fmt.Println(recover())
	}()

	panic("PANIC!!")

}
