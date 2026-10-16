package sqlMethods

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func OpenDB() (*sql.DB, error) {
	username := "root"
	password := "vijay692"
	host := "127.0.0.1"
	port := "3307"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql", username, password, host, port)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func Util() {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := CreateTable(db); err != nil {
		log.Fatal("table creation failed: ", err)
	}

	target := 6
	details, err := FindRows(db, target)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("❌ MySQL returned: No details found with ID %d\n", target)
			return
		} else {
			log.Fatal("Query execution failed: ", err)
		}
	}
	fmt.Println("\n🔎 --- MySQL details Found ---")
	fmt.Printf("ID: %d | Name: %s | Age: %d | Email: %s | Phone: %s | Available: %t\n",
		details.ID, details.Name, details.Age, details.Email, details.PhoneNumber, details.Available)

	target1 := 3
	updatedRows, err := UpdateAvailability(db, target1, false)
	if err != nil {
		log.Fatal("update failed: ", err)
	}
	if updatedRows == 0 {
		fmt.Println("no rows are updated")
	} else {
		fmt.Println("rows are effected :", updatedRows)
	}

	target4 := 1
	deletedRows, err := DeleteDetails(db, target4)
	if err != nil {
		log.Fatal("delete failed: ", err)
	}
	if deletedRows == 0 {
		fmt.Println("no rows are deleted")
	} else {
		fmt.Println("rows are effected :", deletedRows)
	}
}
