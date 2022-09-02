package main

import "fmt"

func main() {
	names := []string{"shahrooz", "Ali"}

	fmt.Println("names", names)
	fmt.Printf("%T \n", names)
	fmt.Println("names", len(names))
	fmt.Println("names", cap(names))
	//An array has a fixed size.
	//A slice, on the other hand, is a dynamically-sized
	//flexible view into the elements of an array.
	//In practice, slices are much more common than arrays.
	// The basic difference between a
	//slice and an array is that a slice is a reference
	// to a contiguous segment of an array.
	// Unlike an array, which is a value-type,
	//slice is a reference type.
	//توضیح بالا یعنی اینکه توی کپی کردن عملا نیازی نیست
	//پوینتر بذاریم
}
