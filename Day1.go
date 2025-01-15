package main

import "fmt"

func main() {
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}

/*
package main lets the Go compiler know that we want this code to compile and run as a standalone program, as opposed to being a library that's imported by other programs.
import "fmt" imports the fmt (formatting) package from the standard library. It allows us to use fmt.Println to print to the console.
func main() defines the main function, the entry point for a Go program.
*/
