package main

import (
	"fmt"
	"math/big"
)

const (
	parallelism = 100
	precision   = 332160
	iterations  = 100000
)

func nilakantha(ch chan *big.Float, istart, istop int) {
	factor := big.NewFloat(0.0).SetPrec(precision)
	mant := big.NewFloat(0.0).SetPrec(precision)

	num := big.NewFloat(4.0)
	den := big.NewFloat(0.0)

	for i := istart; i < istop; i++ {
		x := float64((i + 1) * 2)
		den.SetFloat64(float64(x * (x + 1) * (x + 2)))
		factor.Quo(num, den)

		if i&1 == 0 {
			mant.Add(mant, factor)
		} else {
			mant.Sub(mant, factor)
		}
	}
	ch <- mant
}

func main() {
	// Round up number of iterations to desired parallelism.
	iter := iterations
	if iterations%parallelism != 0 {
		iter += parallelism - (iterations % parallelism)
	}

	results := make(chan *big.Float, parallelism)

	pi := big.NewFloat(3.0).SetPrec(precision)

	chunk := iter / parallelism
	for i := 0; i < parallelism; i++ {
		istart := i * chunk
		istop := istart + chunk
		go nilakantha(results, istart, istop)
	}
	for i := 0; i < parallelism; i++ {
		r := <-results
		pi.Add(pi, r)
	}
	fmt.Println(pi.Text('f', 99991))
}
