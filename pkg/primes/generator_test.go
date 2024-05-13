package primes

import (
	"context"
	"github.com/mvlipka/bhe-code-exercise/pkg/primes/calculators"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"strings"
	"testing"
)

func TestEratosthenesCalculator_GetPrimeAtIndex(t *testing.T) {
	generator := NewGenerator()
	calculator := calculators.NewEratosthenesCalculator()

	result, err := generator.GetPrimeAtIndex(context.Background(), 0, calculator)
	require.NoError(t, err)
	assert.Equal(t, int64(2), result)

	result, err = generator.GetPrimeAtIndex(context.Background(), 19, calculator)
	require.NoError(t, err)
	assert.Equal(t, int64(71), result)

	result, err = generator.GetPrimeAtIndex(context.Background(), 15, calculator)
	require.NoError(t, err)
	assert.Equal(t, int64(53), result)

	result, err = generator.GetPrimeAtIndex(context.Background(), 99, calculator)
	require.NoError(t, err)
	assert.Equal(t, int64(541), result)

	result, err = generator.GetPrimeAtIndex(context.Background(), 500, calculator)
	require.NoError(t, err)
	assert.Equal(t, int64(3581), result)

	result, err = generator.GetPrimeAtIndex(context.Background(), 986, calculator)
	require.NoError(t, err)
	assert.Equal(t, int64(7793), result)

	result, err = generator.GetPrimeAtIndex(context.Background(), 2000, calculator)
	require.NoError(t, err)
	assert.Equal(t, int64(17393), result)

	result, err = generator.GetPrimeAtIndex(context.Background(), 1000000, calculator)
	require.NoError(t, err)
	assert.Equal(t, int64(15485867), result)
}

func TestEratosthenesCalculator_GetPrimeAtIndexTimeout(t *testing.T) {
	generator := NewGenerator()
	calculator := calculators.NewEratosthenesCalculator()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 0)
	defer cancelFunc()

	_, err := generator.GetPrimeAtIndex(ctx, 2000, calculator)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "context timeout exceeded")
}

func TestEratosthenesCalculator_GetPrimeAtIndexInputError(t *testing.T) {
	mockCalculator := &calculators.MockCalculator{
		TestGeneratePrimesInRange: func(start, end int64) ([]int64, error) {
			require.FailNow(t, "this should have been reached")
			return nil, nil
		},
	}
	generator := NewGenerator()

	result, err := generator.GetPrimeAtIndex(context.Background(), int64(-1), mockCalculator)
	require.Error(t, err)
	assert.Equal(t, result, int64(-1))
	assert.Contains(t, err.Error(), "index must be a positive number")
}

func TestEratosthenesCalculator_GetPrimeAtIndexPrimeGenerationError(t *testing.T) {
	mockCalculator := &calculators.MockCalculator{
		TestGeneratePrimesInRange: func(start, end int64) ([]int64, error) {
			return []int64{}, nil
		},
	}

	generator := NewGenerator()

	result, err := generator.GetPrimeAtIndex(context.Background(), int64(100), mockCalculator)
	require.Error(t, err)
	assert.Equal(t, int64(-1), result)
	assert.Contains(t, err.Error(), "no more primes available")

}

func FuzzEratosthenesCalculator_GetPrimeAtIndex(f *testing.F) {
	generator := NewGenerator()
	calculator := calculators.NewEratosthenesCalculator()

	f.Fuzz(func(t *testing.T, n int64) {
		result, err := generator.GetPrimeAtIndex(context.Background(), n, calculator)
		if err != nil {
			if !strings.Contains(err.Error(), "index must be a positive number") {
				t.Errorf("unexpected error: %v", err)
			}
		} else {
			if !big.NewInt(result).ProbablyPrime(0) {
				t.Errorf("the generators produced a non-prime number at index %d", n)
			}
		}
	})
}
