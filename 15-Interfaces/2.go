package main

import "fmt"

func main() {
	myInt := IntCounter(0)
	for i := 0; i < 10; i++ {
		fmt.Println(myInt.Increment())
	}
}

type Incrementer interface {
	Increment() int
}
type IntCounter int

func (ic *IntCounter) Increment() int {
	// fmt.Println(ic)
	*ic++
	return (int(*ic))
}
