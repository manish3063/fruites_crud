package main

import (
	"fruites_crud/connect"
	"fruites_crud/connect/db"
)

func main() {

	db.CreateDbConn()
	db.Redisconn()
	connect.Connect()

	//redis connection..

}
