package apihandler

import (
	"Raspi-TV-Control/pkg/constants/enums/powerstatus"
	"Raspi-TV-Control/pkg/handler/cechandler"
	"Raspi-TV-Control/pkg/system"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/activedevices", ActiveDevicesHandler).Methods("GET")
	r.HandleFunc("/powerstatus", PowerStatusHandler).Methods("GET")
	r.HandleFunc("/powerstatus/{status}", SetPowerStatusHandler).Methods("POST")
	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(system.GetInfo())
	w.Write(data)
}

func ActiveDevicesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(cechandler.GetActiveDevices())
	w.Write(data)
}

func PowerStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(map[string]string{"powerstatus": cechandler.GetPowerStatus()})
	w.Write(data)
}

func SetPowerStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]
	if status != powerstatus.POWERSTATUS_ON && status != powerstatus.POWERSTATUS_STANDBY {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid power status"))
		return
	}
	err := cechandler.SetPowerStatus(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := json.Marshal(map[string]string{"status": "error", "error": err.Error()})
		w.Write(data)
		return
	}
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(map[string]string{"status": "ok"})
	w.Write(data)
}
