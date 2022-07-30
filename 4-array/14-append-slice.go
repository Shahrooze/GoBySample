package main

import "fmt"

func main() {
	primes := make([]int, 1, 10000000)
	primes = append(primes, 100)
	primes = append(primes, 200)
	primes = append(primes, 300)
	fmt.Println(primes)
	fmt.Println(cap(primes))
	fmt.Println(len(primes))
}
