package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "LOGIN PAGE")
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You are in the register page, and you can create a new account ")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
