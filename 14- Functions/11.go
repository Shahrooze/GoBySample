// Go program to illustrate
// use of an anonymous function
package main

import "fmt"

func main() {

	//در اینجا ما یک تابع را به عنوان یک متغیر ذخیره کرده ایم
	value := func() {
		fmt.Println("Welcome! to Go by Haj Shahrooz")
	}
	//اینجا تابع ذخیره شده را با پرانتر با و پرانتز بسته
	//فراخوانی کرده ایم
	value()
}
