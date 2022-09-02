package main

import "fmt"

func main() {
	printAges(10, 22, 34, 3434, 54, 454, 54, 54, 5)
}
func printAges(ages ...int) int {
	sum := 0
	for _, v := range ages {
		sum += v
	}
	fmt.Println(sum)
	return sum
}
