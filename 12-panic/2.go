package main

import "fmt"

func main() {
	fmt.Println("First")
	defer fmt.Println("Before")
	panic("This is panic test")
	defer fmt.Println("After")
	fmt.Println("Last")
}
