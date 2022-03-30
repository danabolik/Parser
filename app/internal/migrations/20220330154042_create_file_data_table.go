package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateNewDataTable, downCreateNewDataTable)
}

func upCreateNewDataTable(tx *sql.Tx) error {
	tx.Exec("CREATE TABLE `file_data` (" +
		"id int not null," +
		"file_name varchar(255) NOT NULL," +
		"extension varchar(4) NOT NULL," +
		"`data` text NOT NULL," +
		"parsing_date  timestamp NOT NULL," +
		"PRIMARY KEY (id)" +
		")")
	return nil
}

func downCreateNewDataTable(tx *sql.Tx) error {
	tx.Exec("DROP TABLE file_data")
	return nil
}
