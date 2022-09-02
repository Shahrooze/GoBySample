package main

import "fmt"

func main() {
	number := 19
	// guess := 30
	guess := 1
	if guess < number {
		fmt.Println("too low")
	}
	if guess > number {
		fmt.Println("too high")
	}
	if guess == number {
		fmt.Println("You Correct")
	}
}
