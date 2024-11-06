package tasks

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	//custom libraries from the project
	database "go-task-manager/handlers/database"
	utilis "go-task-manager/utilis"
)

func Create(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//read request body
		tasks, err := utilis.ReadRequestBody(r)
		if utilis.DisplayErrors(w, "Failed to read request body", err, http.StatusInternalServerError) {
			return
		}

		//Extract the category from the request payload
		category := utilis.ExtractString(tasks, "category")

		switch category {

		case "tasks":

			//Define variables from the request payload
			userId := utilis.Extractfloat(tasks, "user_id")
			title := utilis.ExtractString(tasks, "title")
			description := utilis.ExtractString(tasks, "description")
			dueDate := utilis.ExtractString(tasks, "due_date")
			priority := utilis.ExtractString(tasks, "priority")

			//insert the record to the tasks table
			taskId, err := database.InsertToTasks(db, userId, title, description, priority, dueDate)
			if utilis.DisplayErrors(w, "Failed to insert tasks to the database", err, http.StatusInternalServerError) {
				return
			}

			//Extract the subtasks payload and insert to sub tasks table
			subtasks := tasks["subtasks"].([]interface{})
			subtaskIds := extractSubtasks(db, taskId, subtasks)

			//Define response payload
			response := map[string]interface{}{
				"task_id":     taskId,
				"subtask_ids": subtaskIds,
				"message":     "task was created successfully",
			}

			//write success response
			utilis.CreateSuccessResponse(w, response, http.StatusCreated)

		case "subtasks":
			//define individual variables
			taskId := utilis.Extractfloat(tasks, "task_id")
			description := utilis.ExtractString(tasks, "description")
			priority := utilis.ExtractString(tasks, "priority")

			//insert into subtask table
			subtaskId, err := database.InsertToSubTasks(db, taskId, description, priority)
			if utilis.DisplayErrors(w, "Failed to insert subtask to table", err, http.StatusInternalServerError) {
				return
			}

			//Define response payload
			response := map[string]interface{}{
				"subtask_id": subtaskId,
				"message":    "subtask created successfully",
			}

			//write a success response
			utilis.CreateSuccessResponse(w, response, http.StatusCreated)

		}
	}
}

func extractSubtasks(db *sqlx.DB, taskId int, subtasks []interface{}) []int {

	var subtaskIds []int

	//loop through the outer map string interface
	for _, v := range subtasks {
		subtask, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		//Define the payload from the subtask payload
		description := utilis.ExtractString(subtask, "description")
		priority := utilis.ExtractString(subtask, "priority")

		//insert each record into the subtasks table
		subtaskId, err := database.InsertToSubTasks(db, taskId, description, priority)
		if utilis.LogErrors("Failed to insert to db", err) {
			return nil
		}

		subtaskIds = append(subtaskIds, subtaskId)

	}

	return subtaskIds
}
