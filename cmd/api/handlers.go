package main

import (
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
	_ = app.writeJson(w, http.StatusOK, payload)

}
func (app *application) movies(w http.ResponseWriter, r *http.Request) {
	res, err := app.DB.Movies()
	if err != nil {
		app.errorJson(w, err)
		return
	}
	_ = app.writeJson(w, http.StatusOK, res)
}
