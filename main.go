package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type msg struct {
	Message string
}

func main() {
	http.HandleFunc("/", test)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func test(w http.ResponseWriter, r *http.Request) {
	a := msg{"instagram api clone by sai"}
	m, _ := json.Marshal(a)
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}
