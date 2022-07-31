package main

import "fmt"

func main() {
	//ساختار کلید مقداری دارد
	//امکان تعریف کلید تکراری وجود ندارد.
	//تربیت اینکه به در زمان فراخوانی  مشخص نیست
	fullNames := map[string]int{
		"Shahroozjafari": 32,
		"AliJafari":      65,
	}
	fmt.Println(fullNames)
}
