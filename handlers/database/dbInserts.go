package database

import (
	"go-task-manager/utilis"

	"github.com/jmoiron/sqlx"
)

func InsertToUsers(db *sqlx.DB, username string, email string, password string) (int, error) {

	query := `INSERT INTO users(username, email, password) VALUES (?, ?, ?)`

	result, err := db.Exec(query, username, email, password)
	if utilis.LogErrors("Failed to insert record into users table", err) {
		return 0, err
	}

	userId, err := result.LastInsertId()
	if utilis.LogErrors("Failed to retrieve user id", err) {
		return 0, err
	}

	return int(userId), nil
}

func InsertToTasks(db *sqlx.DB, userId int, title, description, priority, dueDate string) (int, error) {

	query := `INSERT INTO tasks (user_id, title, description, priority, due_date, status) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, userId, title, description, priority, dueDate, "todo")
	if utilis.LogErrors("Failed to insert record into tasks table", err) {
		return 0, err
	}

	taskId, err := result.LastInsertId()
	if utilis.LogErrors("Failed to retrieve task id", err) {
		return 0, err
	}

	return int(taskId), nil

}

func InsertToSubTasks(db *sqlx.DB, taskId int, description string, priority string) (int, error) {
	query := `INSERT INTO sub_tasks(task_id, description, status, priority) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(query, taskId, description, "todo", priority)

	if utilis.LogErrors("Failed to insert record into subtasks table", err) {
		return 0, err
	}

	subtaskId, err := result.LastInsertId()

	if utilis.LogErrors("Failed to insert record into subtasks table", err) {
		return 0, err
	}
	return int(subtaskId), err
}
