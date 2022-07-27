package main

import (
	c "fruites_crud/connect"
	"fruites_crud/connect/db"
)

func main() {

	db.CreateDbConn()
	c.Connect()

}
