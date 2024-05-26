package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Server starting on port 7654...")
	if err := http.ListenAndServe(":7654", nil); err != nil {
		panic(err)
	}
}
