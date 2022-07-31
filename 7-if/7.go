package main

import "fmt"

func main() {
	number := 19
	// guess := 30
	guess := 19
	if guess > 100 || guess < 0 {
		fmt.Println("out of range")
	} else {
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
}
