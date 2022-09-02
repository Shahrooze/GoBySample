package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var wg = sync.WaitGroup{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("sayHello", counter)
	wg.Done()
}
func increment() {
	counter++
	wg.Done()
}
