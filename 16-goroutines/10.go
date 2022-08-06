package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func main() {
	fmt.Println("thread ", runtime.GOMAXPROCS(-1))
}
