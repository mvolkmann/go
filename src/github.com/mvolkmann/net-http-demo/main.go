package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func (handler myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	var handler myHandler
	http.ListenAndServe("localhost:1234", handler)
}
