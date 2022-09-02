package main

import "fmt"

func main() {
	number := 19
	guess := 101
	if guess >= 100 {
		fmt.Println("most less than 100")
	} else if guess <= 1 {
		fmt.Println("most grater than zero")
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
