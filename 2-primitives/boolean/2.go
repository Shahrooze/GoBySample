package main

import "fmt"

func main() {
	//result of logical
	var isSomthing bool = 100 > 1000
	fmt.Printf("%v %T \n", isSomthing, isSomthing)
	var isSomthing1 bool = 100 < 1000
	fmt.Printf("%v %T \n", isSomthing1, isSomthing1)
	i := 1 == 3
	fmt.Printf("%v %T \n", i, i)
	z := 1 < 3
	fmt.Printf("%v %T \n", z, z)

}
