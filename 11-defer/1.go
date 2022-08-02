package main

import "fmt"

func main() {
	//defer
	// باعث می شود که آن خط دقیقا قبل از
	//تمام شدن تابع اجرا شود.
	fmt.Println("First")
	defer fmt.Println("Second")
	fmt.Println("Last")
}
