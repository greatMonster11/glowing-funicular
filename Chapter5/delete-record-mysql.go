package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	CONN_HOST        = "localhost"
	CONN_PORT        = "8080"
	DRIVER_NAME      = "mysql"
	DATA_SOURCE_NAME = "root:password@/mydb"
)

var db *sql.DB
var connectionError error

func init() {
	db, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if connectionError != nil {
		log.Fatal("error connecting to database ::", connectionError)
		return
	}
}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Print("going to delete record in database for name :: ", name[0])
		stmt, err := db.Prepare("DELETE from employee where name=?")
		if err != nil {
			log.Printf("error occurred while preparing query :: ", err)
			return
		}

		result, err := stmt.Exec(name[0])

		if err != nil {
			log.Printf("error occurred while executing :: ", err)
			return
		}

		rowsAffected, err := result.RowsAffected()
		fmt.Fprintf(w, "Number of rows deleted in database are :: %d", rowsAffected)
	} else {
		fmt.Fprint(w, "Error occurred while deleting record in database for name %s", name[0])
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("employee/delete", deleteRecord).Methods("DELETE")

	defer db.Close()
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server:", err)
		return
	}
}
