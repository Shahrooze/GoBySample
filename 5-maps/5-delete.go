package main

import (
	"fmt"
)

func main() {
	fullNames := make(map[string]int)
	fullNames = map[string]int{
		"Shahroozjafari": 32,
		"AliJafari":      65,
	}
	delete(fullNames, "Shahroozjafari")
	fmt.Println(fullNames)
}
