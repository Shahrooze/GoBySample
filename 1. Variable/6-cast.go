package main

import "fmt"

func main() {
	var i float32 = 1
	//درصد تی در واقع داره تایپ متغیر رو نشون میده
	fmt.Printf("%v %T \n", i, i)
	var z int = int(i)
	fmt.Printf("%v %T \n", z, z)
}
