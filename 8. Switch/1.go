package main

import "fmt"

func main() {
	age := 10
	switch age {
	case 1:
		{
			fmt.Println("value is 1")
		}
	case 2:
		{
			fmt.Println("value is 2")
		}
	default:
		{
			fmt.Println("value is other")
		}
	}
}
