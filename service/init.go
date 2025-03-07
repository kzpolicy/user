package service

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

func init() {
	// DB接続
	db, err := sql.Open("mysql", "moizumi:jamyuki0210@tcp(localhost:3306)/ambitious?parseTime=true")
	if err != nil {
		log.Fatalf("Cannot connect database: %v", err)
	}
	boil.SetDB(db)
}