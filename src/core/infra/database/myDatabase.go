package mydatabase

import (
	"os"

	"github.com/jmoiron/sqlx"
)

// DbInit() データベースに接続する。
func DbInit() *sqlx.DB {

	DBMS := os.Getenv("mysql")
	USER := os.Getenv("user")
	PASS := os.Getenv("password")
	DBNAME := os.Getenv("mydb")
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"

	conn, err := sqlx.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}

	//接続オブジェクトのポインタを返す。
	return conn
}
