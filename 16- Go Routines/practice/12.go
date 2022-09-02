package main

import "fmt"

func main() {
	c := make(chan string)
	go func() {
		c <- "Hello shahrooz"
	}()
	msg := <-c
	fmt.Println(msg)
}
