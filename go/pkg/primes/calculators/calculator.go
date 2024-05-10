package calculators

import "context"

// Calculator a general prime calculator interface
type Calculator interface {
	GetPrimeAtIndex(ctx context.Context, n int64) (int64, error)
}
