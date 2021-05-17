package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db := connectToPostgres()

	defer db.Close()

	err := db.Ping()
	CheckError(err)

	startServer()
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

/*
 * Utility Functions
 */

func connectToPostgres() *sql.DB {
	host := os.Getenv("POSTGRES_HOST")
	port, e := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	if e != nil {
		fmt.Println("Invalid port, using 5432 instead")
		port = 5432
	}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	return db
}

func startServer() {
	fmt.Println("Starting server...")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products/{id}", DeleteHandler).Methods("DELETE")
	r.HandleFunc("/products", StockDisplayHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8090", r)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
