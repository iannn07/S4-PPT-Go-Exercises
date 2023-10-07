package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open database connection
	// The database is called test
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ppt_ses8")

	// The error couldn't be read by the system

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
		defer db.Close()
	}

	insert(db)

	fmt.Printf("\nDo you want to delete the data => 2501983105? (Y/N) ")
	var confirm string
	fmt.Scan(&confirm)
	if confirm == "Y" {
		delete(db)
	} else {
		db.Close()
	}
}

func insert(db *sql.DB) {
	//perform a db.Query insert
	insert, err := db.Query("INSERT INTO user VALUES ('2501983105', 'Ian', 'admin@gmail.com', '089687126528')")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func delete(db *sql.DB) {
	del, err := db.Query("DELETE FROM user WHERE nim=2501983105")

	if err != nil {
		panic(err.Error())
	} else {
		del.Close()
	}
}
