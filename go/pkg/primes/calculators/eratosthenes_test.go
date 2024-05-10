package calculators

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestEratosthenesCalculator_GetPrimeAtIndex(t *testing.T) {
	sieve := NewEratosthenesCalculator()

	result, err := sieve.GetPrimeAtIndex(context.Background(), 0)
	require.NoError(t, err)
	assert.Equal(t, int64(2), result)

	result, err = sieve.GetPrimeAtIndex(context.Background(), 19)
	require.NoError(t, err)
	assert.Equal(t, int64(71), result)

	result, err = sieve.GetPrimeAtIndex(context.Background(), 15)
	require.NoError(t, err)
	assert.Equal(t, int64(53), result)

	result, err = sieve.GetPrimeAtIndex(context.Background(), 99)
	require.NoError(t, err)
	assert.Equal(t, int64(541), result)

	result, err = sieve.GetPrimeAtIndex(context.Background(), 500)
	require.NoError(t, err)
	assert.Equal(t, int64(3581), result)

	result, err = sieve.GetPrimeAtIndex(context.Background(), 986)
	require.NoError(t, err)
	assert.Equal(t, int64(7793), result)

	result, err = sieve.GetPrimeAtIndex(context.Background(), 2000)
	require.NoError(t, err)
	assert.Equal(t, int64(17393), result)

	result, err = sieve.GetPrimeAtIndex(context.Background(), 100001)
	require.NoError(t, err)
	assert.Equal(t, int64(1299743), result)

	result, err = sieve.GetPrimeAtIndex(context.Background(), 1000000)
	require.NoError(t, err)
	assert.Equal(t, int64(15485867), result)
}

func TestEratosthenesCalculator_GetPrimeAtIndexTimeout(t *testing.T) {
	sieve := NewEratosthenesCalculator()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 0)
	defer cancelFunc()

	_, err := sieve.GetPrimeAtIndex(ctx, 2000)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "context timeout exceeded")
}

func FuzzEratosthenesCalculator_GetPrimeAtIndex(f *testing.F) {
	sieve := NewEratosthenesCalculator()

	f.Fuzz(func(t *testing.T, n int64) {
		result, err := sieve.GetPrimeAtIndex(context.Background(), n)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !big.NewInt(result).ProbablyPrime(0) {
			t.Errorf("the sieves produced a non-prime number at index %d", n)
		}
	})
}
