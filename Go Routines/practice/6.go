package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		count("Sheep")
		wg.Done()
	}()
	wg.Wait()
}

func count(str string) {
	for i := 0; i < 100; i++ {
		fmt.Println(i, str)
		time.Sleep(time.Millisecond * 500)
	}
}
