package main

import "fmt"

func main() {
	age := 2
	switch age {
	case 1, 2, 3:
		fmt.Println("value is 1-2")
	default:
		{
			fmt.Println("value is other")
		}
	}
}
