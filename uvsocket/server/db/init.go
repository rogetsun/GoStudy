package db

import (
	"GoStudy/uvsocket/serverConf"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func init() {
	fmt.Println("init db connection ... ")
	var err error
	Conn, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			serverConf.DB_USER, serverConf.DB_PASSWORD, serverConf.DB_SERVER, serverConf.DB_PORT, serverConf.DB_SCHAME))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//Conn.Ping()
	Conn.SetMaxOpenConns(2000)
	Conn.SetMaxIdleConns(1000)
	fmt.Println(Conn)
}

func DBTest() error {
	return Conn.Ping()
}
