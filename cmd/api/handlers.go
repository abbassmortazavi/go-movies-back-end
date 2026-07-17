package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) Hello(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "ok",
		Message: "Hello World",
		Version: "1.0.0",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	out, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	w.Write(out)

}
