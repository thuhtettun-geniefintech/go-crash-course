package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("server starting..")
	http.ListenAndServe(":3000", nil)
}
