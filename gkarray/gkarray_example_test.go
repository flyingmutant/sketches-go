package gk_test

import (
	"fmt"
	"math/rand"

	gk "github.com/DataDog/sketches-go/gkarray"
)

func Example() {
	rand.Seed(1234)

	sketch := gk.NewDefaultGKArray()

	for i := 0; i < 500; i++ {
		v := rand.NormFloat64()
		sketch.Add(v)
	}

	anotherSketch := gk.NewDefaultGKArray()
	for i := 0; i < 500; i++ {
		v := rand.NormFloat64()
		anotherSketch.Add(v)
	}
	sketch.Merge(anotherSketch)

	fmt.Println(quantiles(sketch))
	fmt.Println(quantiles(anotherSketch))
	// Output:
	// [0.05752841178130713 0.7615197295175491 1.3493584822458902 3.0655851632106974]
	// [0.15958772878548855 0.862994039450761 1.4702742458010833 3.0282161483661967]
}

func quantiles(sketch *gk.GKArray) []float64 {
	qs := []float64{0.5, 0.75, 0.9, 1}
	quantiles := make([]float64, len(qs))
	for i, q := range qs {
		quantiles[i] = sketch.Quantile(q)
	}
	return quantiles
}
