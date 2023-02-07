package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DATABASE *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", `root:root@tcp(127.0.0.1:8889)/users`)
	if err != nil {
		fmt.Println(err.Error())
	}
	DATABASE = db
}
