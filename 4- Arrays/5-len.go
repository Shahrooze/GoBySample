package main

import "fmt"

func main() {
	var grades [3]int
	grades[0] = 1
	grades[2] = 1
	fmt.Println(len(grades))
}
