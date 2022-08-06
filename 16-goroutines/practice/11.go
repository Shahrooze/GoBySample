package main

import "fmt"

func main() {
	c := make(chan string)
	c <- "Hello shahrooz"
	msg := <-c
	fmt.Println(msg)
}
