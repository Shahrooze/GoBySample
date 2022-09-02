package main

import (
	"fmt"
	"log"
	"net/http"
)

//http://localhost:8080/
//http://localhost:8080/hello
//http://localhost:8080/form.html
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "./static/form.html")
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "فرم ارسالی معتبر نمیباشد  %w", err)
			return
		}
		fmt.Fprintf(w, "درخواست با موفقیت دریافت شد")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "name= %s\n", name)
		fmt.Fprintf(w, "address=%s\n", address)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "یافت نشد", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "متد ارسالی معتبر نیست", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "hello\n")
}
