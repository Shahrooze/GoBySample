package main

import "fmt"

func main() {
	var age interface{} = 10
	switch age.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	default:
		{
			fmt.Println("value is other")
		}
	}
}
