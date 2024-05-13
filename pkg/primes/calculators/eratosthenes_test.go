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

func BenchmarkEratosthenesCalculator_GeneratePrimesInRange(b *testing.B) {
	calculator := EratosthenesCalculator{}

	for n := 0; n < b.N; n++ {
		calculator.GeneratePrimesInRange(2, 1000000)
	}
}
