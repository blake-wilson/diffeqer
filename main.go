package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/blake-wilson/diffeq/methods"
)

var timestep, minrange, maxrange float64

func simpleFunc(params ...float64) float64 {
	return 3 * params[0] * params[0]
}

func simpleFuncDeriv(params ...float64) float64 {
	return 6 * params[0]
}

func evaluateDiffeq() *DiffeqData {
	times, estimates := diffeq.Taylor(simpleFunc, 1, minrange, maxrange, timestep, simpleFuncDeriv)
	return &DiffeqData{
		Time:      times,
		Estimates: estimates,
	}
}

func maxRangeHandler(w http.ResponseWriter, r *http.Request) {
	maxrange, _ = strconv.ParseFloat(r.URL.Path[len(MaxRangeConstant):], 64)

	data := evaluateDiffeq()
	writeResponse(w, data)
}

func timeStepHandler(w http.ResponseWriter, r *http.Request) {

	// get number of iterations from http request
	responseIterations, _ := strconv.ParseFloat(r.URL.Path[len(TimeStepConstant):], 64)
	timestep = maxrange / responseIterations

	data := evaluateDiffeq()
	writeResponse(w, data)
}

func writeResponse(w http.ResponseWriter, data *DiffeqData) {
	json, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func init() {
	timestep = 0.01
	minrange = 0
	maxrange = 10
	http.HandleFunc(TimeStepConstant, timeStepHandler)
	http.HandleFunc(MaxRangeConstant, maxRangeHandler)
	http.ListenAndServe(":8080", nil)
}
