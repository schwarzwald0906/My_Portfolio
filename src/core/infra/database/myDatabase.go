package myDatabase

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DbInit() データベースに接続する。
func DbInit() *sqlx.DB {

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
