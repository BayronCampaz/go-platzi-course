package main

import (
	"fmt"
	pk "go-platzi-course/mypackage"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	myCar := pk.Car{Brand: "Tesla", Year: 2020}
	fmt.Println(myCar)
	pk.PrintMessage("This works")

	//Pointers
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := pc{ram: 8, disk: 516, brand: "HP"}
	fmt.Println(myPc)
	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRAM()

	fmt.Println(myPc)
	myPc.duplicateRAM()

	fmt.Println(myPc)
}
