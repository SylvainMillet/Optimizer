package main

import (
	"database/sql"
	"fmt"
	"os"
)

func main() {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// example config :
	// user:password@tcp(127.0.0.1:3306)/database

	conn, err := sql.Open("mysql", "MarcMo:C3Yq2iQchGEIscOS@tcp(127.0.0.1:3306)/optimizer")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// use your own select statement
	// this is just an example statement
	statement, err := conn.Prepare("SELECT nom FROM patient LIMIT 10")

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

	conn.Close()
}
