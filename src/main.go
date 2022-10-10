package main

import (
	"fmt"
	pk "go-platzi-course/mypackage"
)

func main() {
	myCar := pk.Car{Brand: "Tesla", Year: 2020}
	fmt.Println(myCar)
	pk.PrintMessage("This works")
}
