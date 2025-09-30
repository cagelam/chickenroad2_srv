package randx

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/seehuhn/mt19937"
)

var globalRand = New(&lockedSource{src: mt19937.New()})

// var globalRand = New(&lockedSource{src: rand.NewSource(1).(rngSource)})

func init() {
	b := new(big.Int).SetUint64(uint64(time.Now().UTC().UnixNano() / int64(os.Getpid())))
	sd, _ := crand.Int(crand.Reader, b)
	seed64 := splitmix64(sd.Uint64())
	seed := int64(seed64.Next() >> 1)
	globalRand.Seed(seed)
	rand.Seed(seed)
}

func RandomSeed(state int64) int64 {
	seed64 := splitmix64(state)
	return int64(seed64.Next() >> 1)
}

func Seed(seed int64) { globalRand.Seed(seed) }

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64
// from the default Source.
func Int63() int64 { return globalRand.Int63() }

// Uint32 returns a pseudo-random 32-bit value as a uint32
// from the default Source.
func Uint32() uint32 { return globalRand.Uint32() }

// Uint64 returns a pseudo-random 64-bit value as a uint64
// from the default Source.
func Uint64() uint64 { return globalRand.Uint64() }

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32
// from the default Source.
func Int31() int32 { return globalRand.Int31() }

// Int returns a non-negative pseudo-random int from the default Source.
func Int() int { return globalRand.Int() }

// Int63n returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n <= 0.
func Int63n(n int64) int64 { return globalRand.Int63n(n) }

// Int31n returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n <= 0.
func Int31n(n int32) int32 { return globalRand.Int31n(n) }

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n <= 0.
func Intn(n int) int { return globalRand.Intn(n) }

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0)
// from the default Source.
func Float64() float64 { return globalRand.Float64() }

// Float32 returns, as a float32, a pseudo-random number in the half-open interval [0.0,1.0)
// from the default Source.
func Float32() float32 { return globalRand.Float32() }

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers
// in the half-open interval [0,n) from the default Source.
func Perm(n int) []int { return globalRand.Perm(n) }

// Shuffle pseudo-randomizes the order of elements using the default Source.
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func Shuffle(n int, swap func(i, j int)) { globalRand.Shuffle(n, swap) }

// Read generates len(p) random bytes from the default Source and
// writes them into p. It always returns len(p) and a nil error.
// Read, unlike the Rand.Read method, is safe for concurrent use.
func Read(p []byte) (n int, err error) { return globalRand.Read(p) }

// NormFloat64 returns a normally distributed float64 in the range
// [-math.MaxFloat64, +math.MaxFloat64] with
// standard normal distribution (mean = 0, stddev = 1)
// from the default Source.
// To produce a different normal distribution, callers can
// adjust the output using:
//
//	sample = NormFloat64() * desiredStdDev + desiredMean
func NormFloat64() float64 { return globalRand.NormFloat64() }

// ExpFloat64 returns an exponentially distributed float64 in the range
// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1) from the default Source.
// To produce a distribution with a different rate parameter,
// callers can adjust the output using:
//
//	sample = ExpFloat64() / desiredRateParameter
func ExpFloat64() float64 { return globalRand.ExpFloat64() }

//-------------------------------------------------------------------------------------

// Int63r generates pseudo random int64 between low and high.
//
//	Input:
//	 low  -- lower limit
//	 high -- upper limit
//	Output:
//	 random int64
func Int63r(low, high int64) int64 { return globalRand.Int63r(low, high) }

// Int63s generates pseudo random integers between low and high.
//
//	Input:
//	 low    -- lower limit
//	 high   -- upper limit
//	Output:
//	 values -- slice to be filled with len(values) numbers
func Int63s(values []int64, low, high int64) { globalRand.Int63s(values, low, high) }

// Int63Shuffle - shuffles a slice of integers
func Int63Shuffle(values []int64) { globalRand.Int63Shuffle(values) }

// Uint32 is int range generates pseudo random uint32 between low and high.
//
//	Input:
//	 low  -- lower limit
//	 high -- upper limit
//	Output:
//	 random uint32
func Uint32r(low, high uint32) uint32 { return globalRand.Uint32r(low, high) }

// Uint32s generates pseudo random integers between low and high.
//
//	Input:
//	 low    -- lower limit
//	 high   -- upper limit
//	Output:
//	 values -- slice to be filled with len(values) numbers
func Uint32s(values []uint32, low, high uint32) { globalRand.Uint32s(values, low, high) }

// Uint32Shuffle shuffles a slice of integers
func Uint32Shuffle(values []uint32) { globalRand.Uint32Shuffle(values) }

// Uint64r generates pseudo random uint64 between low and high.
//
//	Input:
//	 low  -- lower limit
//	 high -- upper limit
//	Output:
//	 random uint64
func Uint64r(low, high uint64) uint64 { return globalRand.Uint64r(low, high) }

// Uint64s generates pseudo random integers between low and high.
//
//	Input:
//	 low    -- lower limit
//	 high   -- upper limit
//	Output:
//	 values -- slice to be filled with len(values) numbers
func Uint64s(values []uint64, low, high uint64) { globalRand.Uint64s(values, low, high) }

// Uint64Shuffle - shuffles a slice of integers
func Uint64Shuffle(values []uint64) { globalRand.Uint64Shuffle(values) }

// Int31r is int range generates pseudo random int32 between low and high.
//
//	Input:
//	 low  -- lower limit
//	 high -- upper limit
//	Output:
//	 random int32
func Int31r(low, high int32) int32 { return globalRand.Int31r(low, high) }

// Int31s generates pseudo random integers between low and high.
//
//	Input:
//	 low    -- lower limit
//	 high   -- upper limit
//	Output:
//	 values -- slice to be filled with len(values) numbers
func Int31s(values []int32, low, high int32) { globalRand.Int31s(values, low, high) }

// Int31Shuffle - shuffles a slice of integers
func Int31Shuffle(values []int32) { globalRand.Int31Shuffle(values) }

// Intr is int range generates pseudo random integer between low and high.
//
//	Input:
//	 low  -- lower limit
//	 high -- upper limit
//	Output:
//	 random integer
func Intr(low, high int) int { return globalRand.Intr(low, high) }

// Ints generates pseudo random integers between low and high.
//
//	Input:
//	 low    -- lower limit
//	 high   -- upper limit
//	Output:
//	 values -- slice to be filled with len(values) numbers
func Ints(values []int, low, high int) { globalRand.Ints(values, low, high) }

// IntShuffle shuffles a slice of integers
func IntShuffle(values []int) { globalRand.IntShuffle(values) }

// Float64r generates a pseudo random real number between low and high; i.e. in [low, right)
//
//	Input:
//	 low  -- lower limit (closed)
//	 high -- upper limit (open)
//	Output:
//	 random float64
func Float64r(low, high float64) float64 { return globalRand.Float64r(low, high) }

// Float64s generates pseudo random real numbers between low and high; i.e. in [low, right)
//
//	Input:
//	 low  -- lower limit (closed)
//	 high -- upper limit (open)
//	Output:
//	 values -- slice to be filled with len(values) numbers
func Float64s(values []float64, low, high float64) { globalRand.Float64s(values, low, high) }

// Float64Shuffle shuffles a slice of float point numbers
func Float64Shuffle(values []float64) { globalRand.Float64Shuffle(values) }

// Float32r generates a pseudo random real number between low and high; i.e. in [low, right)
//
//	Input:
//	 low  -- lower limit (closed)
//	 high -- upper limit (open)
//	Output:
//	 random float32
func Float32r(low, high float32) float32 { return globalRand.Float32r(low, high) }

// Float32s generates pseudo random real numbers between low and high; i.e. in [low, right)
//
//	Input:
//	 low  -- lower limit (closed)
//	 high -- upper limit (open)
//	Output:
//	 values -- slice to be filled with len(values) numbers
func Float32s(values []float32, low, high float32) { globalRand.Float32s(values, low, high) }

// Float32Shuffle shuffles a slice of float point numbers
func Float32Shuffle(values []float32) { globalRand.Float32Shuffle(values) }

// FlipCoin generates a Bernoulli variable; throw a coin with probability p
func FlipCoin(p float64) bool { return globalRand.FlipCoin(p) }

// ------------------------------------------------------------------------

type rngSource interface {
	rand.Source
	Uint64() uint64
}

type lockedSource struct {
	lk  sync.Mutex
	src rngSource
	// src rand.Source
}

func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

func (r *lockedSource) Uint64() (n uint64) {
	r.lk.Lock()
	n = r.src.Uint64()
	r.lk.Unlock()
	return
}

// // seedPos implements Seed for a lockedSource without a race condition.
// func (r *lockedSource) seedPos(seed int64, readPos *int8) {
// 	r.lk.Lock()
// 	r.src.Seed(seed)
// 	*readPos = 0
// 	r.lk.Unlock()
// }

// // read implements Read for a lockedSource without a race condition.
// func (r *lockedSource) read(p []byte, readVal *int64, readPos *int8) (n int, err error) {
// 	r.lk.Lock()
// 	n, err = read(p, r.src, readVal, readPos)
// 	r.lk.Unlock()
// 	return
// }
