package main

import mTypes "github.com/blake-wilson/diffeq/types"

type DiffeqResponse struct {
	Time      []float64 `json:"time"`
	Estimates []float64 `json:"estimates"`
	Error     string    `json:"string"`
}

func splitEvalSlice(points []*mTypes.EvalPoint) *DiffeqResponse {
	times := make([]float64, len(points), len(points))
	estimates := make([]float64, len(points), len(points))

	for i, point := range points {
		times[i] = point.Time
		estimates[i] = point.Value
	}
	return &DiffeqResponse{
		Time:      times,
		Estimates: estimates,
	}
}
