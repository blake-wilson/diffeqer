package main

import (
	"encoding/json"
	"net/http"

	"github.com/blake-wilson/diffeq/methods"
)

func simpleFunc(params ...float64) float64 {
	return 3 * params[0] * params[0]
}

func simpleFuncDeriv(params ...float64) float64 {
	return 6 * params[0]
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, estimates := diffeq.Taylor(simpleFunc, 1, 0, 4, 0.01, simpleFuncDeriv)
	json, err := json.Marshal(estimates)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func init() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
