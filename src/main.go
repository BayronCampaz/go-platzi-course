package main

import (
	"fmt"
	"math"
)

func getSquareArea(side float64) float64 {
	return side * side
}

func getCircleArea(radius float64) float64 {
	return math.Pow(radius, 2) * math.Pi
}

func getTrapezoidArea(shortBase, longBase, height float64) float64 {
	return ((shortBase + longBase) / 2) * height
}

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

	fmt.Println(getSquareArea(5))
	fmt.Println(getCircleArea(7))
	fmt.Println(getTrapezoidArea(4, 8, 6))
}
