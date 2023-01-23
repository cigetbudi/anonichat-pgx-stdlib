package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("pgx", "postgres://yzwzagyo:9-rTL5cubS2j7ILQVPonpjJPhvYkg0Yk@topsy.db.elephantsql.com/yzwzagyo")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Terhubung dengan database")

}
