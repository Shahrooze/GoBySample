package main

import "fmt"

func main() {
	grades := [...]int{100, 200, 400, 500, 600}
	fmt.Printf("%v %T \n ", grades, grades)
}
