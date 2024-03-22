package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("hello world")

	sql.Open("postgres", "host=localhost port=5432 user=postgres password=hasibuan")
}