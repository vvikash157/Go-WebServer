package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("server started at port: 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "wrong values passed", http.StatusBadRequest)
		return
	}
	fmt.Println("POST request is successful!!!")

	Name := r.FormValue("name")
	Address := r.FormValue("address")

	fmt.Fprintf(w, "Name= %s \n", Name)
	fmt.Fprintf(w, "Address= %s \n", Address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "hello World")

}
