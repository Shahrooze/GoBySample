package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("Sheep", c)
	//این صرفا
	//Syntactic sugar
	//قبلی هست
	//In computer science, syntactic sugar is syntax
	// within a programming language that is designed
	//to make things easier to read or to express.
	for msg := range c {
		fmt.Println(msg)
	}
}

func count(str string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- str
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
