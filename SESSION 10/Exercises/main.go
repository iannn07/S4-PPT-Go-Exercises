package main

import (
	"encoding/json"
	"io"
	"net/http"
	USR "user"
)

func check_error(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func SelectAll(w http.ResponseWriter, r *http.Request) {
	doc := USR.SelectAll()
	data, err := json.Marshal(doc)
	check_error(err)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(data))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/doctor", SelectAll)
	http.ListenAndServe(":5050", mux)

}
