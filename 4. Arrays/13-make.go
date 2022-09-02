package main

import "fmt"

func main() {
	primes := make([]int, 1, 10000000)
	fmt.Println(primes)
	fmt.Println(cap(primes))
	fmt.Println(len(primes))
	//cap tells you the capacity of the underlying array
	// len tells you how many items are in the array.
}
