package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "sayHello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Bye"
	time.Sleep(1000 * time.Microsecond)
}
