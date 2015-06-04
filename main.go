package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/blake-wilson/diffeq/methods"
	"github.com/blake-wilson/exparser"
)

type diffeqOpts struct {
	// fields used for numerical method
	Timestep    float64 `json:"timestep"`
	InitialTime float64 `json:"initial_time"`
	FinalTime   float64 `json:"final_time"`

	// string representation of diffeq.
	Expression string `json:"expression"`

	// numerical method to use for evaluation
	Method string `json:"method"`
}

var opts diffeqOpts

func evaluateDiffeq(expression string) *DiffeqData {
	var times, estimates []float64

	exprLower := strings.ToLower(expression)
	function, err := exparser.EvalExpression(exprLower)
	if err != nil {
		// error parsing expression
		return nil
	}

	switch opts.Method {
	case "taylor":
		times, estimates = diffeq.Taylor(function, 1, opts.InitialTime, opts.FinalTime, opts.Timestep, function)
	case "euler":
		times, estimates = diffeq.Euler(function, 1, opts.InitialTime, opts.FinalTime, opts.Timestep)
	}
	return &DiffeqData{
		Time:      times,
		Estimates: estimates,
	}
}
func requestHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&opts)
	if err != nil {
		// Could not decode JSON
		return
	}
	fmt.Printf("OPTS: %v\n\n", opts)
	response := evaluateDiffeq(opts.Expression)

	fmt.Printf("resonse:\n\n", response)
	writeResponse(w, response)
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
	opts.Timestep = 0.01
	opts.InitialTime = 0
	opts.FinalTime = 10
	http.HandleFunc("/", requestHandler)
}

func main() {
	http.ListenAndServe(":8080", nil)
}
