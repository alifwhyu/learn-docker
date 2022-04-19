package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Alif Wahyu Widi Adrian")
		fmt.Fprintf(w, "205150601111022")
	})

	http.ListenAndServe(":8080", nil)
}
