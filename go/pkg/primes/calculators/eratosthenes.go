package calculators

import (
	"errors"
	"math"
)

var _ Calculator = &EratosthenesCalculator{}

// EratosthenesCalculator is used to calculate prime numbers via the Sieve of Eratosthenes algorithm
// The calculator stores a cache of calculated prime numbers. Accessing an already existing number will pull from the cache
// More information on the algorithm can be found on Wikipedia: https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
type EratosthenesCalculator struct {
}

// NewEratosthenesCalculator returns a new EratosthenesCalculator
func NewEratosthenesCalculator() *EratosthenesCalculator {
	return &EratosthenesCalculator{}
}

// GeneratePrimesInRange given a range delineated by start & end, gather all prime numbers within the range
func (e *EratosthenesCalculator) GeneratePrimesInRange(start, end int64) ([]int64, error) {
	if start <= 0 {
		return nil, errors.New("error start of range must be positive")
	}

	if start > end {
		return nil, errors.New("error start of range must be less than end of range")
	}

	return e.segmentedSieve(start, end), nil
}

// sieve computes all prime numbers up to max using the classic Sieve of Eratosthenes algorithm
func (e *EratosthenesCalculator) sieve(max int64) []int64 {
	isPrime := make([]bool, max+1)
	for i := int64(2); i <= max; i++ {
		isPrime[i] = true
	}

	primeNumberEstimate := int64(float64(max) / math.Log(float64(max)))
	primes := make([]int64, 0, primeNumberEstimate)
	for i := int64(2); i <= max; i++ {
		if isPrime[i] {
			primes = append(primes, i)
			for j := i * i; j <= max; j += i {
				isPrime[j] = false
			}
		}
	}

	return primes
}

// segmentedSieve a slight tweak to the classic Sieve of Eratosthenes in which we can reduce the memory usage when generating for large primes
func (e *EratosthenesCalculator) segmentedSieve(low, high int64) []int64 {
	root := int64(math.Sqrt(float64(high)))
	primes := e.sieve(root)

	segmentSize := high - low + 1
	isPrime := make([]bool, segmentSize)
	for i := range isPrime {
		isPrime[i] = true
	}

	for _, prime := range primes {
		smallestMultiple := prime * prime
		if smallestMultiple < low {
			smallestMultiple = ((low + prime - 1) / prime) * prime
		}

		for multiple := smallestMultiple; multiple <= high; multiple += prime {
			isPrime[multiple-low] = false
		}
	}

	primeNumbers := make([]int64, 0)
	for i, prime := range isPrime {
		if prime {
			primeNumbers = append(primeNumbers, low+int64(i))
		}
	}
	return primeNumbers
}
