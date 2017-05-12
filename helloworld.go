package main

import (
	"fmt"
	"net/http"
)

func Say(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	http.HandleFunc("/", Say)
	http.ListenAndServe(":8080", nil)
}
