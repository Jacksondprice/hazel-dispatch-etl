package handlers

import (
	"encoding/json"
	"net/http"
)

type Grade struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListGrades(w http.ResponseWriter, _ *http.Request) {
	items := []Grade{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
