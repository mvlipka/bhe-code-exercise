package calculators

type MockCalculator struct {
	TestGeneratePrimesInRange func(start, end int64) ([]int64, error)
}

func (m *MockCalculator) GeneratePrimesInRange(start, end int64) ([]int64, error) {
	return m.TestGeneratePrimesInRange(start, end)
}
