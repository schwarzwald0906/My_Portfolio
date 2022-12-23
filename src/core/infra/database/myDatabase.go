package mydatabase

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

// DbInit() データベースに接続する。
func DbInit() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("USER")
	PASS := os.Getenv("PASSWORD")
	DBNAME := os.Getenv("DBNAME")
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"

	conn, err := sqlx.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}

	//接続オブジェクトのポインタを返す。
	return conn
}
