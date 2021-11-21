package main

import (
	"math/rand"
	"reflect"
	"runtime"
)

type Rules struct {
	feeder Feed
}

type Feed func(*Cell, float64)

func GetRules(feedString string) Rules {
	return Rules{getFeeder(feedString)}
}

func getFeeder(s string) Feed {
	switch s {
	case "constant":
		return constantFeeder
	case "random":
		return randomFeeder
	case "trait":
		return trait1Feeder
	default:
		debugf("Invalid feed parameter '%s', using constant", s)
		return constantFeeder
	}

}

func constantFeeder(cell *Cell, rate float64) {
	cell.food += rate
}

func randomFeeder(cell *Cell, rate float64) {
	food := rand.Float64() * rate
	cell.food += food
}

func trait1Feeder(cell *Cell, rate float64) {
	m := 1.0
	if cell.trait == 1 && cell.age > 50 {
		m = 1.5
	}
	randomFeeder(cell, rate*m)
}

func (r *Rules) AsString() string {
	return "Feeder function: " + GetFunctionName(r.feeder)
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
