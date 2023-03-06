package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:RLGH3qjs!!@tcp(localhost:3306)/addrbook")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	conn := &db_conn{db: db}

	conn.insert(address{"hio", 65})

	result := conn.init()
	for _, v := range result {
		fmt.Printf("name: %s, age: %d\n", v.s_name, v.i_age)
	}
}
