package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "sayHello"
	go func() {
		//با اینکه در یک استک دیگر اجرا می شود
		//ولی بای این حال به متغیر اسکوب قبلی دسترسی
		//دارد.
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Microsecond)
}
