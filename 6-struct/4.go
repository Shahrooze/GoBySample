package main

import "fmt"

func main() {
	//anonymous struct
	aPerson := struct {
		name   string
		family string
	}{name: "shahrooz", family: "Jafari"}
	fmt.Println(aPerson)
}
