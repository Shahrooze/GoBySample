package main

import "fmt"

func main() {
	s := "this is test"
	s[2] = "G"
	fmt.Printf("%v %T \n", s[0], s[1])
}
