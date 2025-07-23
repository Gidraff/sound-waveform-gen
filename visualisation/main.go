package main

import (
	"fmt"
	"math"
)

const period = 2 * math.Pi

func main() {
	sampleRate := 100
	noteA4 := 440.0 // Frequency of A4 note in Hz

	for i := 0; i < sampleRate; i++ {
		sineVal := math.Sin(period * noteA4 * float64(i) / float64(sampleRate))
		fmt.Printf("%f\n", sineVal)
	}
}
