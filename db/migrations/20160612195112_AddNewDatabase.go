package main

import (
	"database/sql"
	"fmt"
)

// Up is executed when this migration is applied
func Up_20160612195112(txn *sql.Tx) {

	_, err := txn.Query(`CREATE TABLE tasks (
		id BIGINT NOT NULL AUTO_INCREMENT,
		title CHAR(255),
		description CHAR(255),
		priority INT NOT NULL DEFAULT '2',
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP,
		completed_at DATETIME,
		is_deleted TINYINT(1),
		is_completed TINYINT(1),
		PRIMARY KEY(id)
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
