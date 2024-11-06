package utilis

import (
	"log"
)

func ExtractString(payload map[string]interface{}, name string) string {

	//check if key exists in the request payload
	value, exists := payload[name]
	if !exists {
		log.Printf("Key %s not found in payload", name)
		return ""
	}

	//check if the value and type assertions are okay from the payload
	response, ok := value.(string)
	if !ok {
		log.Printf("Failed to read %v from request payload", response)
		return ""
	}
	return response
}

func Extractfloat(payload map[string]interface{}, name string) int {
	//check if key exists in the request payload
	value, exists := payload[name]
	if !exists {
		log.Printf("Key %s not found in payload", name)
		return 0
	}

	//check if the value and type assertions are okay from the payload
	response, ok := value.(float64)
	if !ok {
		log.Printf("Failed to read %v from request payload", response)
		return 0
	}
	return int(response)

}

// func extractInterface(payload interface{}, name string) string {
// 	//check if key exists in the request payload
// 	value, exists := payload.name
// 	if !exists {
// 		log.Printf("Key %s not found in payload", name)
// 		return ""
// 	}

// 	//check if the value and type assertions are okay from the payload
// 	response, ok := value.(string)
// 	if !ok {
// 		log.Printf("Failed to read %v from request payload", response)
// 		return ""
// 	}
// 	return response
// }
