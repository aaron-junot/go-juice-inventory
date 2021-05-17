package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting server...")
	r := mux.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8090", r)
}
