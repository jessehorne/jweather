package common

import (
	"encoding/json"
	"net/http"
)

func APIResponse(w http.ResponseWriter, status int, data map[string]interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("json goofed"))
		return
	}

	w.WriteHeader(status)
	w.Write(j)
}
