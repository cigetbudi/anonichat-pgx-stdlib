package db

import (
	"anonichat-pgx-stdlib/utils"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB
var DBN *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("pgx", utils.GetEnv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Terhubung dengan database")

}

func InitDBNEon() {
	var err error
	DBN, err = sql.Open("pgx", utils.GetEnv("DBN_URL"))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Terhubung dengan database Neon")
}
