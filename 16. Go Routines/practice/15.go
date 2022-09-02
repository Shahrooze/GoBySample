package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Eevry 500 ml"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "Eevry 2000 s "
			time.Sleep(time.Second * 2000)
		}
	}()

	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}
