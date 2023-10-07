package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// It's similar with struct in its declaration
type BangunDatar interface {
	Area() float64
	Round() float64
}

type InputSquare struct {
	Side    int
	area_s  int
	round_s int
}

type InputRectangle struct {
	Height int
	Length int
}

// Function to calculate the square
func (Input InputSquare) Area() float64 {
	area_square := Input.Side * Input.Side
	as_struct := InputSquare{
		Side:   Input.Side,
		area_s: area_square,
	}
	return float64(as_struct)
}

func (Input InputSquare) Round() float64 {
	round_square := 4 * Input.Side
	return float64(round_square)
}

// Function to calculate a rectangle
func (Input InputRectangle) Area() float64 {
	area_rect := Input.Length * Input.Height
	return float64(area_rect)
}

func (Input InputRectangle) Round() float64 {
	round_rect := (2 * Input.Length) + (2 * Input.Height)
	return float64(round_rect)
}

func square(w http.ResponseWriter, r *http.Request) {
	square_data := InputSquare{
		Side: 10,
	}
	square_area := square_data.Area()
	// square_round := square_data.Round()

	w.Header().Set("Content-Type", "application/json")
	s_area, err := json.Marshal(square_area)
	// s_round, err := json.Marshal(square_round)
	if err != nil {
		fmt.Print(err)
	} else {
		io.WriteString(w, string(s_area))
		// io.WriteString(w, string(s_round))
	}
}

func rectangle(w http.ResponseWriter, r *http.Request) {
	rectangle_data := InputRectangle{
		Length: 20,
		Height: 10,
	}
	rectangle_area := rectangle_data.Area()
	// rectangle_round := rectangle_data.Round()

	w.Header().Set("Content-Type", "application/json")
	r_area, err := json.Marshal(rectangle_area)
	// r_round, err := json.Marshal(rectangle_round)
	if err != nil {
		fmt.Print(err)
	} else {
		io.WriteString(w, string(r_area))
		// io.WriteString(w, string(r_round))
	}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/square", square)
	mux.HandleFunc("/rectangle", rectangle)

	http.ListenAndServe(":5050", mux)
}
