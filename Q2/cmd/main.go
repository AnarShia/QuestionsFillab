package main

import "fmt"

func main() {
	generateOutput(3)
}

func generateOutput(input int) {

	// divide input by 2 recursively until input is less than or equal to 1
	if input > 1 {
		generateOutput(input / 2)
		fmt.Println(input)
	}
}
