package main

import "fmt"

func main() {
	//ساختار کلید مقداری دارد
	//امکان تعریف کلید تکراری وجود ندارد.
	fullNames := map[string]int{
		"Shahrooz jafari": 32,
		"Ali Jafari":      65,
	}
	fmt.Println(fullNames)
}
