package main

import "fmt"

func main() {

	//Constant declaration
	const pi float64 = 3.14
	const pi2 = 3.1415

	//Variable declaration
	base := 10
	var heigth int = 15
	var area int

	fmt.Println(base, heigth, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	const squareBase = 5
	squareArea := squareBase * squareBase

	fmt.Println(squareArea)

}
