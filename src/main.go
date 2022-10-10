package main

import (
	"fmt"
	pk "go-platzi-course/mypackage"
)

func main() {
	square := pk.Square{}
	square.SetBase(4)

	rectangle := pk.Rectangle{}
	rectangle.SetBase(2)
	rectangle.SetHeight(6)

	circle := pk.Circle{}
	circle.SetRadius(5)

	shapes := []pk.Shape{square, rectangle, circle}

	for _, value := range shapes {
		fmt.Println(value.Area())
	}
}
