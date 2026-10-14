package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN environment variable is required")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	creatingTable(db)
}

func creatingTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS PRODUCTS (
	ID INT AUTO_INCREMENT PRIMARY KEY,
	NAME VARCHAR(50) NOT NULL,
	AGE INT NOT NULL,
	EMAIL VARCHAR(50) UNIQUE,
	PHONENUMBER VARCHAR(10),
	AVAILABLE BOOLEAN,
	CREATEDATE TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	//WE NEED USE DB OBJECT TO CREATE TABLE
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🚀 Success! Connected to database and PRODUCTS table created flawlessly!")

}
