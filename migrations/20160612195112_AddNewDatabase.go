package main

import (
	"database/sql"
	"fmt"
)

// Up is executed when this migration is applied
func Up_20160612195112(txn *sql.Tx) {

	_, err := txn.Query(`CREATE TABLE tasks (
		Id BIGINT PRIMARY_KEY,
		Title CHAR(255),
		Description CHAR(255),
		Priority INT,
		CreatedAt DATETIME,
		UpdatedAt DATETIME,
		CompletedAt DATETIME,
		IsDeleted TINYINT,
		IsCompleted TINYINT
		    );`)
	if err != nil {
		fmt.Print(err)
	}
}

// Down is executed when this migration is rolled back
func Down_20160612195112(txn *sql.Tx) {
	_, err := txn.Query("DROP TABLE tasks")
	if err != nil {
		fmt.Print(err)
	}

}
