package main

import "fmt"

func main() {
	//با این کار تا زمانی که چنل به بیشتر از ۲ پیام نرسیده بلاک
	//نمیسه
	c := make(chan string, 2)
	c <- "Hello Shahrooz"
	c <- "Hello Ali"
	msg := <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)
}
