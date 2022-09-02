package main

import "fmt"

func main() {
	grades := [3]int{100, 200, 200}
	for k, v := range grades {
		fmt.Println(k, v)
	}
}
