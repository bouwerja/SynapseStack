package databse

import (
	"backend/utils"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func dbConneciton() *sql.DB {
	dbEnv, err := utils.GetDatabaseDetails()
	if err != nil {
		panic(err)
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("Recovered from panic in database:", err)
		}
	}()

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbEnv.Username,
		dbEnv.Password,
		dbEnv.ServerName,
		dbEnv.Port,
		dbEnv.Database,
	)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database unreachable: ", err)
	}

	return db
}

func InsertScraper() any {
	conn := dbConneciton()
	defer conn.Close()
	stats := conn.Stats()
	return stats.OpenConnections
}
