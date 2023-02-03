package main

import (
	"anonichat-pgx-stdlib/api"
	"anonichat-pgx-stdlib/db"
)

func main() {
	db.InitDB()
	db.InitDBNEon()
	r := api.InitRoutes()
	r.Run()
}
