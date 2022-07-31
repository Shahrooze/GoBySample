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
	fullNames["Shahroozjafari"] = 50
	fmt.Println(fullNames["Shahroozjafari"])
}
