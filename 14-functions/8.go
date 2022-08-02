package main

import "fmt"

func main() {
	fmt.Println(divide(1, 0))
	fmt.Println(divide(6, 3))
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divide by zero")
	}
	return (a / b), nil
}
