package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aaron-suarez/go-juice-inventory/database"
	"github.com/aaron-suarez/go-juice-inventory/entities"
	"github.com/aaron-suarez/go-juice-inventory/util"
	"github.com/gorilla/mux"
)

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
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	util.CheckError(err)

	db := database.GetDbInstance()

	_, err = db.Exec("UPDATE juice SET deleted_at = now() WHERE id IN (SELECT id FROM juice WHERE id = $1 LIMIT 1);", id)
	util.CheckError(err)

	fmt.Printf("Successfully performed soft delete on id %d\n", id)
	w.WriteHeader(http.StatusOK)
}

/*
 * GET /products
 */
func StockDisplayHandler(w http.ResponseWriter, req *http.Request) {
	db := database.GetDbInstance()
	rows, err := db.Query("SELECT * FROM juice WHERE deleted_at IS NULL LIMIT 200;")
	util.CheckError(err)

	defer rows.Close()

	var juiceSlice []entities.Juice

	for rows.Next() {
		var (
			id         int64
			name       string
			expiration time.Time
			deleted_at sql.NullString
		)
		err := rows.Scan(&id, &name, &expiration, &deleted_at)
		util.CheckError(err)
		juiceSlice = append(juiceSlice, entities.Juice{Id: id, Name: name, Expiration: expiration})
	}
	prettyJSON, err := json.MarshalIndent(juiceSlice, "", "    ")
	util.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", prettyJSON)
}
