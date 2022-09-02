package main

import "fmt"

func main() {
	sayMessage("ali", 12)
}
func sayMessage(name string, age int) {
	fmt.Printf("hello %v  %v \n", name, age)
}
