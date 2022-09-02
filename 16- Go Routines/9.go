package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("sayHello", counter)
	m.RUnlock()
	wg.Done()
}
func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
