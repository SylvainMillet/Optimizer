package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/optimizer")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// use your own select statement
	// this is just an example statement
	statement, err := db.Prepare("SELECT nom FROM patient limit 10")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := statement.Query() // execute our select statement

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for rows.Next() {
		var title string
		rows.Scan(&title)
		fmt.Println("Title of tutorial is :", title)
	}

	db.Close()

}
