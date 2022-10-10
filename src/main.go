package main

import (
	"fmt"
	"strings"
)

func main() {

	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Slide methods
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)

	mySlice := []string{"How", "you", "doing"}

	for i, value := range mySlice {
		fmt.Println(i, value)
	}

	word := "kayaK"
	fmt.Println(isPalindrome(word))
}

func isPalindrome(word string) bool {
	result := true
	word = strings.ToLower(word)
	for i := 0; i <= len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			result = false
			break
		}
	}

	return result
}
