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
	//اینجا روی یک مپ جدید کپی کردیم
	newFullNames := fullNames
	//یک آیتم رو از مپ جدید حذف کردیم
	delete(newFullNames, "Shahroozjafari")
	fmt.Println(fullNames)
	fmt.Println(newFullNames)
}
