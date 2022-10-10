package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20
	fmt.Println(m)

	for i, value := range m {
		fmt.Println(i, value)
	}

	value, ok := m["Jose"]
	fmt.Println(value, ok)

	//What happens when we try to access an unexistent key
	value2, ok2 := m["Joseph"]
	fmt.Println(value2, ok2)

}
