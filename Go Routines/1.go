package main

import (
	"fmt"
)

func main() {
	//در خط پایین ما یک ترد سبز ران میکنیم
	//green thread
	//باید دربارش بخونید
	go sayHello()
	//خط پایین رو هم برای این نوشتم که بدون این
	//خط و همچنین با این خط تست کنید کد رو
	//و فرقش رو ببینید
	//time.Sleep(100 * time.Microsecond)
}

func sayHello() {
	fmt.Println("sayHello")
}
