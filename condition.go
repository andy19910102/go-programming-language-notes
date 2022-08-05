package main

import "fmt"

func main() {
	yourAge := 10

	if yourAge >= 18 {
		fmt.Println("You can Vote")
	} else if yourAge >= 16 {
		fmt.Println("You can Drive")
	} else {
		fmt.Println("You can Have Fun")
	}
}
