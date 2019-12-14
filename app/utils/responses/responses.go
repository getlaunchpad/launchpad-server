package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Success(w http.ResponseWriter, statusCode int, message string) {
	response := map[string]interface{}{"success": true, "message": message}
	w.WriteHeader(statusCode)
	Respond(w, response)
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	response := map[string]interface{}{"success": false, "message": message}
	w.WriteHeader(statusCode)
	Respond(w, response)
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Fprintf(w, "Error encoding json: %s", err.Error())
	}
}
