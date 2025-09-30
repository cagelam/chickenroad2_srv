package randx

import (
	"math/rand"

	"github.com/seehuhn/mt19937"
)

type Rand struct {
	// lk   sync.Mutex
	rand *rand.Rand
}

// var _ *rand.Rand = (*Rand)(nil)

// New returns a new gosfmt Rand that uses random values from src
// to generate other random values.
func New(source rand.Source) *Rand {
	return &Rand{
		rand: rand.New(source),
	}
}

func NewMT19937() *Rand {
	return &Rand{
		rand: rand.New(mt19937.New()),
	}
}

// Seed uses the provided seed value to initialize the generator to a deterministic state.
// Seed should not be called concurrently with any other Rand method.
func (r *Rand) Seed(seed int64) { r.rand.Seed(seed) }

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *Rand) Int63() int64 { return r.rand.Int63() }

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Rand) Uint32() uint32 { return r.rand.Uint32() }

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (r *Rand) Uint64() uint64 { return r.rand.Uint64() }

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func (r *Rand) Int31() int32 { return r.rand.Int31() }

// Int returns a non-negative pseudo-random int.
func (r *Rand) Int() int { return r.rand.Int() }

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Int63n(n int64) int64 { return r.rand.Int63n(n) }

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Int31n(n int32) int32 { return r.rand.Int31n(n) }

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Intn(n int) int { return r.rand.Intn(n) }

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float64() float64 { return r.rand.Float64() }

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float32() float32 { return r.rand.Float32() }

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func (r *Rand) Perm(n int) []int { return r.rand.Perm(n) }

// Shuffle pseudo-randomizes the order of elements.
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func (r *Rand) Shuffle(n int, swap func(i, j int)) { r.rand.Shuffle(n, swap) }

// Read generates len(p) random bytes and writes them into p. It
// always returns len(p) and a nil error.
// Read should not be called concurrently with any other Rand method.
func (r *Rand) Read(p []byte) (n int, err error) {
	n, err = r.rand.Read(p)
	return n, err
}

// NormFloat64 returns a normally distributed float64 in the range
// [-math.MaxFloat64, +math.MaxFloat64] with
// standard normal distribution (mean = 0, stddev = 1)
// from the default Source.
// To produce a different normal distribution, callers can
// adjust the output using:
//
//	sample = NormFloat64() * desiredStdDev + desiredMean
func (r *Rand) NormFloat64() float64 { return r.rand.NormFloat64() }

// ExpFloat64 returns an exponentially distributed float64 in the range
// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1) from the default Source.
// To produce a distribution with a different rate parameter,
// callers can adjust the output using:
//
//	sample = ExpFloat64() / desiredRateParameter
func (r *Rand) ExpFloat64() float64 { return r.rand.ExpFloat64() }

// ------------------------------------------------------------------------
// SplitMix64 PRNG
type splitmix64 uint64

// Next returns the next number in the sequence
func (x *splitmix64) Next() uint64 {
	*x += 0x9E3779B97F4A7C15
	z := uint64(*x)
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	z = z ^ (z >> 31)
	return z
}
