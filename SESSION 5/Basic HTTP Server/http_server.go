package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to new server!")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome Students!")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome teachers!")
	})

	http.ListenAndServe(":5050", nil)
}
