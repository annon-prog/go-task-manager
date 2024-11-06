package utilis

import (
	"encoding/json"
	"log"
	"net/http"
)

func DisplayErrors(w http.ResponseWriter, message string, err error, statusCode int) bool {
	if err != nil {
		log.Printf("%s: %v", message, err)
		createErrorResponse(w, message, err, statusCode)
		return true
	}
	return false
}
func LogErrors(message string, err error) bool {
	if err != nil {
		log.Printf("%s: %v", message, err)
		return true
	}
	return false
}

func DisplayBoolErrors(w http.ResponseWriter, message string, check bool, statusCode int) bool {
	if !check {
		log.Printf("%s", message)
		createBoolResponse(w, message, statusCode)

		return true
	}
	return false
}

func LogBoolErrors(message string, check bool) bool {
	if !check {
		log.Printf("%s", message)
		return true
	}
	return false
}

func createErrorResponse(w http.ResponseWriter, message string, err error, statusCode int) {
	//create error response
	errorResponse := map[string]interface{}{
		"message": message,
		"error":   err.Error(),
	}

	//encode error response as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse)
}

func createBoolResponse(w http.ResponseWriter, message string, statusCode int) {
	//create error response
	errorResponse := map[string]interface{}{
		"message": message,
	}

	//encode error response as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse)
}

func CreateSuccessResponse(w http.ResponseWriter, response map[string]interface{}, statusCode int) {

	//write Json Response
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
