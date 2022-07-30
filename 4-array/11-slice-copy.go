package main

import "fmt"

func main() {
	names := []string{"shahrooz", "Ali"}
	newNames := names
	newNames[0] = "Nima"
	newNames[1] = "New Name"
	fmt.Println("names", names)
	fmt.Println("newNames", newNames)
}
