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
	//به وسیله کلید به مقدار آن دسترسی پیدا میکنیم
	fmt.Println(fullNames["Shahroozjafari"])
}
