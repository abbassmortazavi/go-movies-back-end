package main

import (
	"errors"
	"log"
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
	res, err := app.Movie.Movies()
	if err != nil {
		app.errorJson(w, err)
		return
	}
	_ = app.writeJson(w, http.StatusOK, res)
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	//read json payload
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := app.readJson(w, r, &payload)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	//validate user against database
	user, err := app.User.GetUserByEmail(payload.Email)
	if err != nil {
		app.errorJson(w, err, http.StatusNotFound)
		return
	}

	//check password

	//create jwt user
	u := jwtUser{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}

	//generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJson(w, err)
		return
	}
	log.Printf("tokens: %v", tokens)
	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokens.AccessToken))
}
