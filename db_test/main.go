package main // mainパッケージであることを宣言

import (
	// fmtモジュールをインポート

	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal("OpenError: ", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("PingError: ", err)
	}
}