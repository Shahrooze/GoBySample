package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("1")
	panicker()
	fmt.Println("2")

}
func panicker() {
	fmt.Println("panicker calling")
	defer func() {
		// The recover built-in function allows a program to manage behavior of a
		// panicking .
		err := recover()
		if err != nil {
			log.Println("Somthing wrong", err)
			panic(err)
		}
	}()
	panic("Panic App!!")
	fmt.Println("panicker called")

}
