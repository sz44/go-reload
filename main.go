package main

import (
	"fmt"
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "404 not found")
}

// type myHandler string
//
// func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, mh)
// }
//
// var mh myHandler = "hello from myHandler"

// type myHandleFunc func(w http.ResponseWriter, r *http.Request)
//
//	func (mh myHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//		mh(w, r)
//	}
//
//	func helloHandler(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w, "Hello from myHandleFunc")
//	}
func notmain() {
	test()
	fmt.Println("hello goodbye")
	// http.Handle("/h", mh)
	// http.Handle("/h", myHandleFunc(helloHandler))
	http.HandleFunc("/", notFoundHandler)
	fmt.Println("Listening on http://127.0.0.1:3000")
	http.ListenAndServe(":3000", nil)
}
