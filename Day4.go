package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Person struct to hold information about a person
type Person struct {
	Name string
	Age  int
}

// NewPerson creates a new person with a given name and age
func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

// Greet method for the Person struct
func (p *Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// RandomNumber generates a random number between min and max
func RandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// Factorial calculates the factorial of a number
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci generates the nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// PrintArray prints the elements of an array
func PrintArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// BubbleSort sorts an array using the bubble sort algorithm
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Create a new person and greet
	person := NewPerson("Alice", 30)
	person.Greet()

	// Generate a random number between 1 and 100
	randomNum := RandomNumber(1, 100)
	fmt.Printf("Random number: %d\n", randomNum)

	// Calculate the factorial of 5
	fact := Factorial(5)
	fmt.Printf("Factorial of 5: %d\n", fact)

	// Generate the 10th Fibonacci number
	fib := Fibonacci(10)
	fmt.Printf("10th Fibonacci number: %d\n", fib)

	// Create an array and print it
	arr := []int{5, 3, 8, 6, 2, 7, 4, 1}
	fmt.Print("Original array: ")
	PrintArray(arr)

	// Sort the array using bubble sort and print it
	BubbleSort(arr)
	fmt.Print("Sorted array: ")
	PrintArray(arr)
}
