package primes

import (
	"context"
	"errors"
	"fmt"
	"github.com/mvlipka/bhe-code-exercise/pkg/primes/calculators"
)

// Generator generates prime numbers using a specified prime number calculator
// Primes will be cached upon generation for later use, if the prime can not be found
// then the prime cache will become invalidated and re-generated
type Generator struct {
	primes     []int64
	calculator calculators.Calculator
}

func NewGenerator(calculator calculators.Calculator) *Generator {
	return &Generator{
		primes:     make([]int64, 0),
		calculator: calculator,
	}
}

func (g *Generator) GetPrimeAtIndex(ctx context.Context, index int64) (int64, error) {
	// Primes must be positive
	if index < 0 {
		return 0, errors.New("index must be a positive number")
	}

	if g.primes == nil || len(g.primes) == 0 {
		g.primes = []int64{}
	}

	// Return the prime at index if it's cached
	if index < int64(len(g.primes)) {
		return g.primes[index], nil
	}

	// Clear our cache to regenerate if the index was missed
	g.primes = make([]int64, 0)

	doneChan := make(chan error)
	var result int64

	go func() {
		start := int64(2)
		end := start + int64(1000)

		// Generate primes until we have a prime slice with an index that satisfies the caller's index
		for {
			newPrimes, err := g.calculator.GeneratePrimesInRange(start, end)
			if err != nil {
				doneChan <- fmt.Errorf("error generating primes in range: %w", err)
			}

			g.primes = append(g.primes, newPrimes...)

			// We found our prime
			if int64(len(g.primes)) > index {
				result = g.primes[index]
				break
			}

			start = end + 1
			end = start + 1000
		}

		doneChan <- nil
	}()

	// Block until primes are generated or the context timeout has exceeded
	select {
	case <-ctx.Done():
		return -1, errors.New("context timeout exceeded")
	case <-doneChan:
		break
	}

	return result, nil
}
