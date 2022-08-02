package main

import "fmt"

func main() {
	a, err := divide(1, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divide by zero")
	}
	return (a / b), nil
}
