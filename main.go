package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello, World!")
	var db = Db{}
	err := db.ConnectDB()
	if err != nil {
		fmt.Printf("%v", err)
	}
	ScanDirectory("/home/andro", db)
}
