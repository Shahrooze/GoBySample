package main

import (
	"fmt"
	"time"
)

func main() {
	count("Sheep")
	count("Fish")
}

func count(str string) {
	for i := 0; true; i++ {
		fmt.Println(i, str)
		time.Sleep(time.Millisecond * 500)
	}
}
