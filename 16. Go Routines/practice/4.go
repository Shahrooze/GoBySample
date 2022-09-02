package main

import (
	"fmt"
	"time"
)

func main() {
	go count("Sheep")
	go count("Fish")
	time.Sleep(2 * time.Second)
}

func count(str string) {
	for i := 0; true; i++ {
		fmt.Println(i, str)
		time.Sleep(time.Millisecond * 500)
	}
}
