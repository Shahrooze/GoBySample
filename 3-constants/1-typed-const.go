package main

import "fmt"

func main() {
	const i int = 10
	//روی کانست ها نمیشه تغییر داد بعد از مقدار دهی اولیه
	//i=100;
	fmt.Printf("%v %T \n", i, i)
}
