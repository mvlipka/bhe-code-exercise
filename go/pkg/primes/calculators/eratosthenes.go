package calculators

import (
	"context"
	"errors"
	"slices"
)

var _ Calculator = &EratosthenesCalculator{}

// EratosthenesCalculator is used to calculate prime numbers via the Sieve of Eratosthenes algorithm
// The calculator stores a cache of calculated prime numbers. Accessing an already existing number will pull from the cache
// More information on the algorithm can be found on Wikipedia: https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
type EratosthenesCalculator struct {
	primes  []int64
	markers []bool
}

// NewEratosthenesCalculator returns a new EratosthenesCalculator
func NewEratosthenesCalculator() *EratosthenesCalculator {
	return &EratosthenesCalculator{
		primes: []int64{2},
	}
}

// GetPrimeAtIndex returns a prime at a given index where index 0 is the first prime number
// This operation may be expensive if a large index is selected
// ctx may be used to set a timeout or deadline
func (e *EratosthenesCalculator) GetPrimeAtIndex(ctx context.Context, index int64) (int64, error) {

	// Primes must be positive
	if index < 0 {
		return 0, errors.New("index must be a positive number")
	}

	if e.primes == nil {
		e.primes = []int64{2}
	}

	primeNumbers := slices.Clone(e.primes)

	// Return the prime at index if it's cached
	if index < int64(len(primeNumbers)) {
		return primeNumbers[index], nil
	}

	doneChan := make(chan bool)
	var result int64

	go func() {
		start := primeNumbers[len(primeNumbers)-1] + 1
		end := start + 1000

		// Generate primes until we have a prime with an index that satisfies the caller's index
		for {
			newPrimes := e.generatePrimesInRange(start, end)
			primeNumbers = append(primeNumbers, newPrimes...)

			// We found our prime
			if int64(len(primeNumbers)) >= index {
				result = primeNumbers[index]
				break
			}

			start = end
			end += 1000
		}

		doneChan <- true
	}()

	// Block until primes are generated or the context timeout has exceeded
	select {
	case <-ctx.Done():
		return -1, errors.New("context timeout exceeded")
	case <-doneChan:
		break
	}

	e.primes = primeNumbers

	return result, nil
}

func (e *EratosthenesCalculator) generatePrimesInRange(start, end int64) []int64 {
	newPrimes := make([]int64, 0)
	markers := make([]bool, end+1)

	// Iterate through the numbers from start to end
	for i := start; i <= end; i++ {
		if markers[i] == true {
			continue
		}

		// Check if the current number is a multiple of any known primes
		isPrime := true
		for _, prime := range append(e.primes, newPrimes...) {
			if i%prime == 0 {
				isPrime = false
				break
			}
		}

		if isPrime == false {
			markers[i] = true
			continue
		}

		// Mark any multiples of the current prime until the end of the segment
		for j := i * i; j < end; j++ {
			if j%i == 0 {
				markers[j] = true
			}
		}

		newPrimes = append(newPrimes, i)
	}

	return newPrimes
}
