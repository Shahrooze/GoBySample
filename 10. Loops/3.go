package main

import "fmt"

func main() {
	grades := [3]int{100, 200, 200}
	for i := 0; i < len(grades); {
		fmt.Println(grades[i])
		i++
	}
}
