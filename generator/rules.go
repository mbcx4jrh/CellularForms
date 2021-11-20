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

func (r *Rules) AsString() string {
	return "Feeder function: " + GetFunctionName(r.feeder)
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
