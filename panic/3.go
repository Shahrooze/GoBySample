package main

import "log"

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("Somthing wrong", err)
		}
	}()
	panic("This is panic test")
}
