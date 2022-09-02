package main

import "fmt"

const (
	i = iota
	z
	v
)
const (
	dd = iota
	cc
)

func main() {
	fmt.Println(i, v, z)
	fmt.Println(dd, cc)
}
