package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("Sheep", c)
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func count(str string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- str
		time.Sleep(time.Millisecond * 500)
	}
}
