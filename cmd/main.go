package main

import (
	"Raspi-TV-Control/pkg/constants"
	"Raspi-TV-Control/pkg/handler/apihandler"
	"Raspi-TV-Control/pkg/handler/arghandler"
	"Raspi-TV-Control/pkg/handler/cechandler"
	"Raspi-TV-Control/pkg/system"
	"log"
	"net/http"
	"strconv"
)

func main() {
	arghandler.HandleArgs()
	system.Init()
	log.Println("Initializing CEC connection...")
	err := cechandler.Init()
	if err != nil {
		log.Fatalf("Error while initializing CEC handler: %s", err)
	}
	r := apihandler.RegisterRoutes()
	http.Handle("/", r)
	log.Println("Starting server on port " + strconv.Itoa(constants.PORT))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(constants.PORT), nil))
}
