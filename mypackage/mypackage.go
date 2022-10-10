package mypackage

import "fmt"

// Public class
type Car struct {
	Brand string
	Year  int
}

// Private class
type car struct {
	brand string
	year  int
}

//Public function
func PrintMessage(message string) {
	fmt.Println(message)
}

//Private function
func printMessage(message string) {
	fmt.Println(message)
}
