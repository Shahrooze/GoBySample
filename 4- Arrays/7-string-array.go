package main

import "fmt"

func main() {
	names := [5]string{}
	names[0] = "Shahrooz"
	names[1] = "Arash"
	names[2] = "Iman"
	names[0] = "Sajjad"
	fmt.Println(names)
	fmt.Println(names[0])
}
