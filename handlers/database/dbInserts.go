package database

import (
	"go-task-manager/utilis"

	"github.com/jmoiron/sqlx"
)

func InsertToUsers(db *sqlx.DB, username string, email string, password string) (int, error) {

	query := `INSERT INTO users(username, email, password) VALUES ($1, $2, $3) RETURNING id`

	var userId int
	err := db.QueryRow(query, username, email, password).Scan(&userId)
	if utilis.LogErrors("Failed to insert record into users table", err) {
		return 0, err
	}

	return userId, nil
}

func InsertToTasks(db *sqlx.DB, userId int, title, description, priority, dueDate string) (int, error) {

	query := `INSERT INTO tasks (user_id, title, description, priority, due_date, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var taskId int
	err := db.QueryRow(query, userId, title, description, priority, dueDate, "todo").Scan(&taskId)
	if utilis.LogErrors("Failed to insert record into tasks table", err) {
		return 0, err
	}

	return taskId, nil

}

func InsertToSubTasks(db *sqlx.DB, taskId int, description string, priority string) (int, error) {
	query := `INSERT INTO sub_tasks(task_id, description, status, priority) VALUES ($1, $2, $3, $4) RETURNING id`

	var subtaskId int
	err := db.QueryRow(query, taskId, description, "todo", priority).Scan(&subtaskId)

	if utilis.LogErrors("Failed to insert record into subtasks table", err) {
		return 0, err
	}

	return subtaskId, err
}
