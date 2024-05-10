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
func (s *EratosthenesCalculator) GetPrimeAtIndex(ctx context.Context, index int64) (int64, error) {

	// Primes must be positive
	if index < 0 {
		return 0, errors.New("index must be a positive number")
	}

	primeNumbers := slices.Clone(s.primes)

	// Return the prime at index if it's cached
	if index < int64(len(primeNumbers)) {
		return primeNumbers[index], nil
	}

	// TODO: move this to a blocking select to adhere to the timeout passed in by the user
	result, primes := generatePrimesToIndex(primeNumbers, index)
	s.primes = primes

	return result, nil
}

func generatePrimesToIndex(knownPrimes []int64, index int64) (int64, []int64) {
	results := make([]bool, 100000)
	primes := knownPrimes

	// Start at the last known prime number
	start := primes[len(primes)-1]

	// End will dictate the range in which we calculate primes until
	// This will change at the end of each iteration as a way to paginated results
	end := start + 1000

	// Loop until we have at least the index's worth of prime numbers
	for {
		if int64(len(primes)) >= index {
			break
		}

		// Iterate through the current page of numbers
		for i := start; i <= end; i++ {
			if results[i] == true {
				continue
			}

			// We need to re-calculate if a number is prime when adding a new page
			// Take the existing prime numbers and check if the current number is a multiple of the prime
			isPrime := true
			for _, prime := range primes {
				if i%prime == 0 {
					results[i] = true
					isPrime = false
					break
				}
			}

			// The number was found to be a multiple of one of the prime numbers, therefore it is composite
			if isPrime == false {
				continue
			}

			results[i] = true
			primes = append(primes, i)

			// Set all the multiples of the current number, within the range, to composite
			for j := i * i; j < end; j++ {
				if j%i == 0 {
					results[j] = true
				}
			}
		}

		start = end + 1
		end = start + 1000
		results = append(results, make([]bool, 1000)...)
	}

	return primes[index], primes
}
