package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}
	c := append(a, b...)
	fmt.Println(c)
}
