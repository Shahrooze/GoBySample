package main

import "fmt"

func main() {
	s := "this is test"
	fmt.Printf("%v %T \n", s, s)
	fmt.Printf("%v %T \n", s[0], s[1])

	s1 := 't'
	fmt.Printf("%v %T \n", s1, s1)

	s3 := []rune(s)
	fmt.Printf("%v %T \n", s3, s3)

}
