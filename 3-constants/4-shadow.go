package main

import "fmt"

const i uint32 = 10000

func main() {
	const i int = 10
	//در این حالت هم اسکوپ داخلی برنده هست یعنی داخل تابع مین
	fmt.Printf("%v %T \n", i, i)
}
