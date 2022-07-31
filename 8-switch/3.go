package main

import "fmt"

func main() {
	age := 2
	switch {
	case age <= 10:
		fmt.Println("age less than 10")
	case age >= 1:
		fmt.Println("age bigger than 1")
	default:
		{
			fmt.Println("value is other")
		}
	}
}
