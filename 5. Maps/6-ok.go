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
	name, ok := fullNames["ss"]
	fmt.Println(name, ok)
	name, ok = fullNames["Shahroozjafari"]
	fmt.Println(name, ok)
}
