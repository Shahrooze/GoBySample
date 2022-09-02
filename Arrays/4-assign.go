package main

import "fmt"

func main() {
	var grades [3]int
	grades[0] = 1
	grades[2] = 9
	fmt.Printf("%v %T \n ", grades, grades)
}
