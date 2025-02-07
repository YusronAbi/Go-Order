package configs

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

func InitDB() {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db
	fmt.Println("Database connected successfully")
}
