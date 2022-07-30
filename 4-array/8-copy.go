package main

import "fmt"

func main() {
	names := [...]string{"shahrooz", "Ali"}
	newNames := names
	newNames[0] = "Nima"
	fmt.Println("name", names)
	fmt.Println("newNames", newNames)
	//نکته پرفورمنسی
	//در زبان های دیگر وقتی آرایه را کپی میکنیم
	//عملا یک متغیر ایجاد میکنیم و همان آرایه قبلی اشاره میکند
	//در زبان گو در واقع عملیات کپی تمام المان ها را کپی میکند
	//کد پایین را در مرورگر اجرا کنیم
	// 	var names=["shahrooz","ali"]
	// var newNames=names;
	// newNames[0]="Nima"
	// console.log("names",names)
	// console.log("newNames",newNames)
}
