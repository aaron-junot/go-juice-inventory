package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aaron-suarez/go-juice-inventory/database"
	"github.com/aaron-suarez/go-juice-inventory/routes"
	"github.com/aaron-suarez/go-juice-inventory/util"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db := database.GetDbInstance()

	defer db.Close()

	err := db.Ping()

	// TODO: remove when the db doesn't need to be initialized from scratch by SetUpDb()
	if err != nil {
		fmt.Println("Waiting for DB to come online...")
		time.Sleep(2 * time.Second)
		err = db.Ping()
	}
	util.CheckError(err)

	// TODO: remove this too once there is a stable production database
	database.SetUpDb(db)

	startServer()
}

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/products/{id}", routes.DeleteHandler).Methods("DELETE")
	r.HandleFunc("/products", routes.StockDisplayHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8090", r)
}
