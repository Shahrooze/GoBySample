package main

import "fmt"

func main() {
	var a int = 1
	var b *int = &a
	fmt.Println(a)
	a = 100
	fmt.Println(a, *b)
	*b = 1000
	fmt.Println(a, *b)

}
