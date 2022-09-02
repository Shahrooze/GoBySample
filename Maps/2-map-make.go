package main

import (
	"fmt"
)

func main() {
	fullNames := make(map[string]int)
	fullNames = map[string]int{
		"Shahrooz jafari": 32,
		"Ali Jafari":      65,
	}
	fmt.Println(fullNames)
}
