package apihandler

import (
	"Raspi-TV-Control/pkg/constants"
	"Raspi-TV-Control/pkg/system"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func RegisterRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(constants.PORT), nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(system.GetInfo())
	w.Write(data)
}
