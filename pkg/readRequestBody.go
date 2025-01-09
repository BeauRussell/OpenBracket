package pkg

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ReadRequestBody(r *http.Request) (error, map[string]interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to parse request body: %v", err)
		return err, nil
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Failed to unmarshal request body to data: %v", err)
		return err, nil
	}

	return nil, data
}
