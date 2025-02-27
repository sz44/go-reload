package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
		fmt.Fprintf(w, "the current time is: %v", time.Now())
	})
	fmt.Println("-- listening on http://127.0.0.1:3001")
	http.ListenAndServe(":3001", nil)
}
