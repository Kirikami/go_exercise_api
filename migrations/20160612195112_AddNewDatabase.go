package main

import (
	"database/sql"
	"fmt"
)

// Up is executed when this migration is applied
func Up_20160612195112(txn *sql.Tx) {

	_, err := txn.Query(`CREATE TABLE tasks (
		Id bigint primary key,
		Title char(255),
		Description char(255),
		Priority int,
		CreatedAt datetime,
		UpdatedAt datetime,
		CompletedAt datetime,
		IsDeleted tinyint,
		IsCompleted tinyint
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
