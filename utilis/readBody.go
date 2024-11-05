package utilis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReadRequestBody(r *http.Request) (map[string]interface{}, error) {
	var requestBody map[string]interface{}

	//read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %v", err)
	}

	//unmarshal request body
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %v", err)
	}

	return requestBody, nil
}
