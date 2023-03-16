package main

import (
	"fmt"
	"salaryCalculator"
	"salaryFine"
)

func main() {
	total := salaryCalculator.Calculate(40, 20000)
	fmt.Println(total)
	fine := salaryFine.FineCalculate(5, 2000)
	fmt.Println(fine)
	fmt.Println("Total Salary = ", total-fine)
}
