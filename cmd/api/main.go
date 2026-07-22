package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"backend/pkg"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

const port = 8080

type application struct {
	Domain       string
	Dsn          string
	DB           repository.Repository
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {

	fmt.Println("Hello World")
	var app application
	//set config

	//read the command line
	flag.StringVar(&app.Dsn, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "PostgreSQL dsn")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "Cookie domain")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "verysecret", "JWT secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "JWT issuer")
	flag.StringVar(&app.Domain, "domain", "example.com", "Domain")
	flag.Parse()

	//connect to the database
	log.Println("Listening on port " + strconv.Itoa(port))

	connection, err := app.connect()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: connection}
	defer pkg.GetDB().Close()
	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
		CookieDomain:  app.CookieDomain,
	}
	//start the webserver
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		panic(err)
	}
}
