package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: ""}
	fmt.Println(RMB, symbol[RMB]) // "3"
	fmt.Println(len(symbol))

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	// d := [3]int{1, 2}
	// fmt.Println(a == d) --> Will return an error since Go detect the length of the array
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
