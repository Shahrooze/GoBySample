package main

import "fmt"

func main() {
	cities := map[string]int{
		"Tehran":    21,
		"Kordestan": 87,
	}
	pop, ok := cities["Tehran"]
	if ok {
		fmt.Println(pop)
	}

}
