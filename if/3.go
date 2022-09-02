package main

import "fmt"

func main() {
	cities := map[string]int{
		"Tehran":    21,
		"Kordestan": 87,
	}
	pop, ok := cities["ss"]
	if ok {
		fmt.Println(pop)
	}

}
