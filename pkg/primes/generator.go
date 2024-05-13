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

// NewGenerator returns a new generator
func NewGenerator(calculator calculators.Calculator) *Generator {
	return &Generator{
		primes:     make([]int64, 0),
		calculator: calculator,
	}
}

// GetPrimeAtIndex generates primes until there are enough primes to satisfy the index requirement
// This method caches the resulting primes. If the cache contains the index, it will simply return the prime in cache
// rather than re-calculate.
func (g *Generator) GetPrimeAtIndex(ctx context.Context, index int64) (int64, error) {
	// Index must be a positive number
	if index < 0 {
		return -1, errors.New("index must be a positive number")
	}

	// Check cache for the index and return it if it's found
	if index < int64(len(g.primes)) {
		return g.primes[index], nil
	}

	start := int64(2)

	// Start from our last known prime+1, if a cached prime exists
	if len(g.primes) > 0 {
		start = g.primes[len(g.primes)-1] + 1
	}

	// Generate until we reach the index requirement
	for int64(len(g.primes)) <= index {
		if err := ctx.Err(); err != nil {
			return -1, fmt.Errorf("context timeout exceeded")
		}

		// Generate primes from start to start + 1000
		newPrimes, err := g.calculator.GeneratePrimesInRange(start, start+1000)
		if err != nil {
			return -1, err
		}

		// If we failed to generate new primes, there may be a problem with our range
		if len(newPrimes) == 0 {
			return -1, fmt.Errorf("no more primes available up to %d", start+1000)
		}

		g.primes = append(g.primes, newPrimes...)
		start += 1001
	}

	return g.primes[index], nil
}
