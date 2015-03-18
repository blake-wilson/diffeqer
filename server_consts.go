package main

type Method int32

const (
	MaxRangeConstant   = "maxrange"
	TimeStepConstant   = "timestep"
	MethodNameConstant = "method"
)

const (
	Euler = iota
	Taylor
)
