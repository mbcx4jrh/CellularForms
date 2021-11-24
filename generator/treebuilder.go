package main

import (
	"log"
	"math/rand"

	"github.com/downflux/go-kd/kd"

	"github.com/downflux/go-kd/point"
)

func BuildTree(points []point.P, s int) *kd.T {
	m := len(points)
	n := m / s

	if n < 100 {
		t, _ := kd.New(points)
		return t
	}

	sample := make([]point.P, 0, n)

	for i := 0; i < n; i++ {
		sample = append(sample, points[rand.Intn(m)])
	}

	debugf("Building tree from sample size %d", len(sample))
	tree, err := kd.New(sample)
	if err != nil {
		log.Fatal("Issue building stochastic tree from sample", err)
	}

	for i := range points {
		tree.Insert(points[i])
	}

	return tree
}
