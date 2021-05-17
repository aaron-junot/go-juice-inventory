package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting server...")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products/{id}", DeleteHandler).Methods("DELETE")
	r.HandleFunc("/products", StockDisplayHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8090", r)
}

/*
 * GET /
 */
func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the Juice Inventory\n")

}

/*
 * DELETE /products/{id}
 */
func DeleteHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Println("Successfully deleted", vars["id"])
	fmt.Fprintf(w, "Successfully deleted %s\n", vars["id"])
}

/*
 * GET /products
 */
func StockDisplayHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is the inventory:\n")
}
