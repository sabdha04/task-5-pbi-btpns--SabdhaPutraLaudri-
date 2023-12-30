package main

import (
	"FinalProject/database"
	"FinalProject/routes" 
)

func main() {
	r := routes.SetupRouter() 
	// Database connection
	database.Connect()

	r.Run(":8080")
}
