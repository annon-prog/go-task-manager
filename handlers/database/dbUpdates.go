package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	//custom library from project
	utilis "go-task-manager/utilis"
)

func UpdateValue(db *sqlx.DB, tableName string, columnName string, value interface{}, id int) {
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE id = ?", tableName, columnName)
	_, err := db.Exec(query, value, id)
	if utilis.LogErrors("Failed to update table", err) {
		return
	}

}
