package main

import "fmt"

func main() {
	a := 1
	b := a
	fmt.Println(a)
	a = 2
	fmt.Println(a, b)

}
