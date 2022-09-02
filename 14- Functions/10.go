package main

import "fmt"

func main() {
	//به این حالت تعریف تابع به صورت
	//anonymous function
	//گفته می شود
	//پرانتز باز و بسته کردن باعث می شود
	//تابع اجرا شود
	func() {
		fmt.Println("This anonymous function")
	}()
}
