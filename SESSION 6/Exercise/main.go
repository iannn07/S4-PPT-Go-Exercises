package main

import (
	"Shape2Dimension"
	"Shape3Dimension"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func squareArea(w http.ResponseWriter, r *http.Request) {
	side := 20
	area := Shape2Dimension.SquareArea(side)

	str, err := json.Marshal(area)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Print(err)
	} else {

		io.WriteString(w, string(str))
	}
	//io.WriteString(w, strconv.FormatFloat(area, 'f', -1, 64))
}

func cylVolume(w http.ResponseWriter, r *http.Request) {
	radius := 10
	h := 15

	vol := Shape3Dimension.CylVolume(radius, h)
	w.Header().Set("Content-Type", "application/json")
	str, err := json.Marshal(vol)
	if err != nil {
		fmt.Print(err)
	} else {

		io.WriteString(w, string(str))
	}
}

func input_cyl(w http.ResponseWriter, r *http.Request) {
	radius, _ := strconv.Atoi(r.FormValue("Radius"))
	h, _ := strconv.Atoi(r.FormValue("Height"))

	vol := Shape3Dimension.CylVolume(radius, h)
	w.Header().Set("Content-Type", "application/json")
	str, err := json.Marshal(vol)
	if err != nil {
		fmt.Print(err)
	} else {

		io.WriteString(w, string(str))
	}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/squareArea", squareArea)
	mux.HandleFunc("/cylVolume", cylVolume)
	mux.HandleFunc("/inputCyl", input_cyl)

	http.ListenAndServe(":5050", mux)
}
