package main

import (
	"fmt"
	"net/http"
	"salaryCalculator"
	"salaryFine"
)

func main() {
	calculate := salaryCalculator.Calculate(40, 300000)
	fineCalculate := salaryFine.FineCalculate(5, 15000)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Total salary")
		fmt.Fprint(w, calculate-fineCalculate)
	})

	http.HandleFunc("/calculator", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, calculate)
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fineCalculate)
	})

	http.ListenAndServe(":5050", nil)
}
