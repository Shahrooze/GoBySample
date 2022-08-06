package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "sayHello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Bye"
	time.Sleep(1000 * time.Microsecond)
}
