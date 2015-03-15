package main

import (
	"encoding/json"
	"net/http"

	"github.com/blake-wilson/diffeq/methods"

	"appengine"
)

func simpleFunc(params ...float64) float64 {
	return 3 * params[0] * params[0]
}

func simpleFuncDeriv(params ...float64) float64 {
	return 6 * params[0]
}

func handler(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	times, estimates := diffeq.Taylor(simpleFunc, 1, 0, 4, 0.01, simpleFuncDeriv)
	packedData := DiffeqData{
		Time:      times,
		Estimates: estimates,
	}
	json, err := json.MarshalIndent(packedData, "", "	")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Infof("Packed data is %s", json)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func init() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
