package main

import "fmt"

type Animal struct {
	name string
}
type Bird struct {
	Animal
	canFly bool
}

func main() {
	//گو مفهوم ارث بری را مثل زبان های شی گرا پشتیبانی نمیکند
	//برای شبیه سازی آن عمل میتوانیم از این مدل استفاده کنیم
	aBird := Bird{}
	aBird.canFly = true
	aBird.name = "Kalagh"
	fmt.Println(aBird)
}
