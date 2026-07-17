package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const port = 8080

type application struct {
	Domain string
}

func main() {

	fmt.Println("Hello World")
	var app application
	//set config

	//read the command line

	//connect to the database
	app.Domain = "example.com"
	log.Println("Listening on port " + strconv.Itoa(port))

	//start the webserver
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		panic(err)
	}
}
