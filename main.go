package main

import (
	"statistics/database"
	"statistics/server"
)

func main() {
	error := database.DatabaseInitSession() // Connect to the database
	if error != nil {
		panic("Failed to connect to the database: " + error.Error())
	} else {
		println("Connected to Cassandra cluster successfully")
	}
	//defer database.Session.Close()

	server.Server()

}
