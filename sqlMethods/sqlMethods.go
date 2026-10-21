package sqlMethods

import (
	"database/sql"
	"fmt"
)

type Details struct {
	ID          int  		`json:"id" validate:"-"`
	Name        string		`json:"name" validate:"required"`
	Age         int			`json:"age" validate:"required"`
	Email       string		`json:"email" validate:"required"`
	PhoneNumber string		`json:"phonenumber" validate:"required"`
	Available   bool		`json:"available" validate:"required"`
}

func CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS PRODUCTS (
	ID INT AUTO_INCREMENT PRIMARY KEY,
	NAME VARCHAR(50) NOT NULL,
	AGE INT NOT NULL,
	EMAIL VARCHAR(50),
	PHONENUMBER VARCHAR(10),
	AVAILABLE BOOLEAN,
	CREATEDATE TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	if _, err := db.Exec(query); err != nil {
		return err
	}

	fmt.Println("🚀 Success! Connected to database and PRODUCTS table created flawlessly!")
	return nil
}

func InsertData(db *sql.DB, product Details) (int64, error) {
	query := `INSERT INTO PRODUCTS (NAME, AGE, EMAIL, PHONENUMBER, AVAILABLE)
		VALUES (?, ?, ?, ?, ?)`

	result, err := db.Exec(query, product.Name, product.Age, product.Email, product.PhoneNumber, product.Available)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func FindRows(db *sql.DB, targetID int) (Details, error) {
	var product Details
	err := db.QueryRow(
		"SELECT ID, NAME, AGE, EMAIL, PHONENUMBER, AVAILABLE FROM PRODUCTS WHERE ID = ?",
		targetID,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Age,
		&product.Email,
		&product.PhoneNumber,
		&product.Available,
	)
	if err != nil {
		return Details{}, err
	}

	return product, nil
}

func ReturnAllDetails(db *sql.DB) ([]Details, error) {
	rows, err := db.Query(
		"SELECT ID, NAME, AGE, EMAIL, PHONENUMBER, AVAILABLE FROM PRODUCTS",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Details
	for rows.Next() {
		var product Details
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Age,
			&product.Email,
			&product.PhoneNumber,
			&product.Available,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func UpdateAvailability(db *sql.DB, targetID int, isAvailable bool) (int64, error) {
	result, err := db.Exec(`UPDATE PRODUCTS SET AVAILABLE = ? WHERE ID = ?`, isAvailable, targetID)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func UpdatePhone(db *sql.DB, targetID int, ph string) (int64, error) {
	result, err := db.Exec(`UPDATE PRODUCTS SET PHONENUMBER = ? WHERE ID = ?`, ph, targetID)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteDetails(db *sql.DB, targetID int) (int64, error) {
	result, err := db.Exec(`DELETE FROM PRODUCTS WHERE ID = ?`, targetID)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
