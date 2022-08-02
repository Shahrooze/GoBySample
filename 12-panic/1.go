package main

import "fmt"

func main() {
	a, b := 1, 0
	if b == 0 {
		panic("b is zero")
	}
	c := a / b
	fmt.Println(c)
}
