package main

import "fmt"

func main() {
	printAges(10, 22, 34, 3434, 54, 454, 54, 54, 5)
}

///سعی کنید از این حالت استفاده نکنید
//برای این گذاشتم که با حالت های مختلف آشنا بشید
func printAges(ages ...int) (sum int) {
	for _, v := range ages {
		sum += v
	}
	fmt.Println(sum)
	return
}
