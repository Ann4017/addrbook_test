package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB_info struct {
	s_user     string
	s_pwd      string
	s_url      string
	s_engine   string
	s_database string
}

var DB1 = DB_info{"root", "RLGH3qjs!!", "localhost:3306", "mysql", "addrbook"}

func (db *DB_info) DB_Query(query string) (name string) {
	dataSource := db.s_user + ":" + db.s_pwd + "@tcp(" + db.s_url + ")/" + db.s_database
	conn, err := sql.Open(db.s_engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.QueryRow(query).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
	return name
}
