package main

import "fmt"

func main() {
	a := 10
	var b int8 = 3
	//خطا میدهد و نیاز هست که حتما تایپ رو تبدیل کنیم
	// fmt.Println(a / b)
	fmt.Println(a / int(b))
}
