package tasks

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	//custom libraries from project
	database "go-task-manager/handlers/database"
	utilis "go-task-manager/utilis"
)

func Update(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//read request payload
		tasks, err := utilis.ReadRequestBody(r)
		if utilis.DisplayErrors(w, "Failed to read request body", err, http.StatusInternalServerError) {
			return
		}

		//Define individual variable
		id := utilis.Extractfloat(tasks, "id")
		category := utilis.ExtractString(tasks, "category")
		taskType := utilis.ExtractString(tasks, "type")
		value := utilis.ExtractString(tasks, "value")

		//update the value to relevant table
		database.UpdateValue(db, category, taskType, value, id)

		//define response payload
		response := map[string]interface{}{
			"message": "successfully updated the record",
		}

		//return a success response
		utilis.CreateSuccessResponse(w, response, http.StatusOK)
	}

}
