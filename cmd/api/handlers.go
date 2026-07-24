package main

import (
	"errors"
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
	valid, err := user.PasswordMatches(payload.Password)
	if err != nil || !valid {
		err := app.errorJson(w, errors.New("invalid credentials"), http.StatusForbidden)
		if err != nil {
			return
		}
		return
	}

	//create jwt user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	//generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	err = app.writeJson(w, http.StatusOK, tokens)
	if err != nil {
		return
	}
}
