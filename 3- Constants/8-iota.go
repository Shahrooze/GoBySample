package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	BT
	EB
	ZB
	YB
)

func main() {
	fmt.Println(3 << 1)
	fmt.Println(3 >> 1)
	fileSize := float32(20000000000)
	fmt.Printf("fileSize %.2fGB ", fileSize/GB)
}
