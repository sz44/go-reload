package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "the current time is: %v", time.Now())
	})
	fmt.Println("listening on http://127.0.0.1:3000")
	http.ListenAndServe(":3000", nil)
}
