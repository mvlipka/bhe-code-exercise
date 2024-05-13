package calculators

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEratosthenesCalculator_GeneratePrimesInRange(t *testing.T) {
	calculator := EratosthenesCalculator{}

	result, err := calculator.GeneratePrimesInRange(int64(5), int64(15))
	require.NoError(t, err)
	assert.Equal(t, []int64{5, 7, 11, 13}, result)
}

func TestNewEratosthenesCalculator_GeneratePrimesInRangeInputError(t *testing.T) {
	calculator := EratosthenesCalculator{}

	result, err := calculator.GeneratePrimesInRange(int64(-1), int64(3))
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error start of range must be positive")

	result, err = calculator.GeneratePrimesInRange(int64(5), int64(4))
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error start of range must be less than end of range")
}

func BenchmarkEratosthenesCalculator_GeneratePrimesInRange(b *testing.B) {
	calculator := EratosthenesCalculator{}

	for n := 0; n < b.N; n++ {
		calculator.GeneratePrimesInRange(2, 1000000)
	}
}
