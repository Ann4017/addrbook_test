package main

import (
	"fmt"
)

func main() {
	db := &DB_info{}
	result := db.DB_Query("select name from info")
	fmt.Println(result)
}
