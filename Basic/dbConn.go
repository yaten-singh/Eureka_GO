package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

)

var db *sql.DB

// This function will make a connection to the database only once.
func init() {
	var err error

	connStr := "postgres://bglmueto:2qb_SexEF4-ADc14fLHfuHyums9u73oC@lucky.db.elephantsql.com/bglmueto"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")
}

func main() {
      	var err error

	psqlconn := "postgres://bglmueto:2qb_SexEF4-ADc14fLHfuHyums9u73oC@lucky.db.elephantsql.com/bglmueto"
	
        // open database
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)
     
        // close database
    defer db.Close()
 
        // check db
    err = db.Ping()
    CheckError(err)
 
    fmt.Println("Connected!")
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}
