package main

import "fmt"

func main() {
	presAge := make(map[string]int)

	presAge["TheodoreRoosevelt"] = 42
	presAge["John F. Kennedy"] = 43

	fmt.Println(presAge["TheodoreRoosevelt"])
	fmt.Println(presAge["John F. Kennedy"])
	fmt.Println(len(presAge))

	delete(presAge, "John F. Kennedy")
	fmt.Println(len(presAge))
}
