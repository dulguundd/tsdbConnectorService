package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
	"tsdbConnectorService/dto"
	"tsdbConnectorService/service"
)

type DataHandlers struct {
	service service.DataService
}

func (ch *DataHandlers) GetTodayClearData(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	datas, err := ch.service.GetTodayClearData()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, datas)
		serviceLatencyLogger(start)
	}
}

func (ch *DataHandlers) getTodayClearDataOfDevice(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	vars := mux.Vars(r)
	id := vars["device_id"]

	idInt, _ := strconv.Atoi(id)

	datas, err := ch.service.GetTodayClearDataOfDevice(idInt)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, datas)
		serviceLatencyLogger(start)
	}
}

func (ch *DataHandlers) PostData(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var req dto.NewDataRequest

	vars := mux.Vars(r)
	id := vars["device_id"]
	req.Device_id, _ = strconv.Atoi(id)
	temperature := vars["temp"]
	req.Temp, _ = strconv.ParseFloat(temperature, 64)

	data, err := ch.service.PostData(req)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, data)
		serviceLatencyLogger(start)
	}
}

func (ch *DataHandlers) PostDataBody(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var req dto.NewDataRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		data, err1 := ch.service.PostData(req)
		if err1 != nil {
			writeResponse(w, err1.Code, err1.Message)
		} else {
			writeResponse(w, http.StatusCreated, data)
			serviceLatencyLogger(start)
		}
	}
}
