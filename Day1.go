package main

import "fmt"

// func main() {
// 	numMessagesFromDoris := 72
// 	costPerMessage := .02
// 	totalCost := costPerMessage * float64(numMessagesFromDoris)
// 	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
// }

/*
package main lets the Go compiler know that we want this code to compile and run as a standalone program, as opposed to being a library that's imported by other programs.
import "fmt" imports the fmt (formatting) package from the standard library. It allows us to use fmt.Println to print to the console.
func main() defines the main function, the entry point for a Go program.
*/

//	func main() {
//		averageOpenRate, displayMessage := .23, "is the average open rate of your messages"
//		fmt.Println(averageOpenRate, displayMessage)
//	}

// func main() {
// 	const name = "Saul Goodman"
// 	const openRate = 30.5
// 	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent%s", name, openRate, "\n")
// 	// ?

// 	// don't edit below this line

// 	fmt.Print(msg)
// }

func main() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	// Don't touch above this line

	// ?
	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.1f, Is Subscribed: %v, Message: %s", fname, lname, age, messageRate, isSubscribed, message)

	// Don't touch below this line

	fmt.Println(userLog)
}
