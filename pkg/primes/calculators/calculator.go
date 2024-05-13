package calculators

// Calculator a general prime calculator interface
type Calculator interface {
	GeneratePrimesInRange(start, end int64) ([]int64, error)
}
