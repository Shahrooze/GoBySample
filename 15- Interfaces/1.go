package main

import (
	"fmt"
)

func main() {
	var cw Writer = ConsoleWriter{}
	var fw Writer = FilterWriter{}
	cw.Write([]byte("This is test "))
	fw.Write([]byte("This is test "))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}
type FilterWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
func (cw FilterWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
