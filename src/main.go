package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {

	fmt.Println("Adios")

	var wg sync.WaitGroup

	wg.Add(1)
	go say("world", &wg)

	wg.Wait()

	go func(text string) { // Funciona anonima
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}
