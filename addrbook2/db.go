package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type address struct {
	s_name string
	i_age  int
}

type db_conn struct {
	db *sql.DB
}

func (conn *db_conn) insert(address address) {
	query := "insert into info (name,age) values (?,?)"
	_, err := conn.db.Exec(query, address.s_name, address.i_age)
	if err != nil {
		log.Fatal(err)
	}
}

func (conn *db_conn) delete(name string) {
	query := "delete from info where name = ? "
	_, err := conn.db.Exec(query, name)
	if err != nil {
		log.Fatal(err)
	}
}

func (conn *db_conn) update(name string, age int) {
	query := "update info set age = ? where name = ?"
	_, err := conn.db.Exec(query, age, name)
	if err != nil {
		log.Fatal(err)
	}
}

func (conn *db_conn) init() []address {
	query := "select * from info"
	rows, err := conn.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []address
	for rows.Next() {
		var a address
		err := rows.Scan(&a.s_name, &a.i_age)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, a)
	}
	return result
}
