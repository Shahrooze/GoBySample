package main

import "fmt"

func main() {
	printAges(10, 22, 34, 343, 54, 454, 54, 54, 5)
}
func printAges(ages ...int) {
	fmt.Println(ages)
}
