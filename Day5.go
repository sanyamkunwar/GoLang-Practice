package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, World!")
	sum := add(3, 4)
	fmt.Println("Sum of 3 and 4 is:", sum)
}
