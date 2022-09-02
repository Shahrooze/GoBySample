package main

import "fmt"

func main() {
	//حذف آخرین آیتم
	a := []int{1, 2, 3, 4, 5}
	c := a[:len(a)-1]
	fmt.Println(c)
	//حذف اولین آیتم
	c = a[1:]
	fmt.Println(c)

}
