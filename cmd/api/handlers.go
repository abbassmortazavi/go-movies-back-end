package main

import (
	"backend/internal/models"
	"encoding/json"
	"net/http"
	"time"
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
func (app *application) movies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	rd, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	newMovie := models.Movie{
		ID:          1,
		Title:       "Movie1",
		Description: "this is a movie",
		ReleaseDate: rd,
		Runtime:     100,
		MPARating:   "1000",
		Image:       "R",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	movies = append(movies, newMovie)
	rd, _ = time.Parse("2007-01-02", "2007-02-10")
	newMovie1 := models.Movie{
		ID:          2,
		Title:       "Movie2",
		Description: "this is a movie2",
		ReleaseDate: rd,
		Runtime:     200,
		MPARating:   "PG-15",
		Image:       "www.testUrl.com",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	movies = append(movies, newMovie1)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	out, err := json.Marshal(movies)
	if err != nil {
		panic(err)
	}
	w.Write(out)
}
