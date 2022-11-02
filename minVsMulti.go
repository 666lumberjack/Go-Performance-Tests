package main

import (
	"math"
	"math/rand"
	"time"
)

// I read recently that in some languages, using min(a, b, ...z) == 0 rather than a series of ORed a == 0, b == 0 results
// in faster execution because after compilation the latter results in a branch in the resulting assembly but the former does not
// This is a test of whether that holds true for Go
func minVsMultiCompare() {
	testCases := make([][3]float64, 1000000)

	for _, tc := range testCases {
		tc[0] = float64(rand.Intn(2))
		tc[1] = float64(rand.Intn(2))
		tc[2] = float64(rand.Intn(2))
	}

	min(testCases)
	multi(testCases)
}

func min(testCases [][3]float64) {
	incrementor := 0
	defer timeTrack(time.Now(), "Comparing Min")
	for _, tc := range testCases {
		if math.Min(tc[0], math.Min(tc[1], tc[2])) == 0 {
			incrementor++
		}
	}
}

func multi(testCases [][3]float64) {
	incrementor := 0
	defer timeTrack(time.Now(), "Comparing all")
	for _, tc := range testCases {
		if tc[0] == 0 || tc[1] == 0 || tc[2] == 0 {
			incrementor++
		}
	}
}
