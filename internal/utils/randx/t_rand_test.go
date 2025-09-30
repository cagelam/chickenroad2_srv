package randx

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestShuffle(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	IntShuffle(arr)
	fmt.Printf("%v\n", arr)
}

func TestRange(t *testing.T) {
	for i := 0; i < 10; i++ {
		val := Intr(0, 10-1)
		fmt.Printf("%v\n", val)
	}
}

func TestSpliteMix64(t *testing.T) {
	seed := splitmix64(12345)
	for i := 0; i < 10; i++ {
		val := seed.Next()
		fmt.Printf("%x | %d\n", val, int64(val))
	}
}

func TestGlobalRand(t *testing.T) {
	var (
		wg     sync.WaitGroup
		loc    sync.Mutex
		values = make([]int32, 10000)
	)
	ts1 := time.Now().UnixNano()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			// 1m
			for n := 0; n < 1000000; n++ {
				v := Int31n(10000)
				loc.Lock()
				values[v]++
				loc.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	ts2 := time.Now().UnixNano()

	for i := 0; i < len(values); i++ {
		fmt.Printf("%5d:%5d\n", i, values[i])
	}
	fmt.Printf("%10s:%5.2f\n", "Avg", float64(SumInt(values))/float64(len(values)))

	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	avg := float64(SumInt(values)) / float64(len(values))
	n := len(values)
	l := n / 4
	m := n / 2
	u := n * 3 / 4
	lower := (values[l] + values[l+1]) / 2
	median := (values[m] + values[m+1]) / 2
	upper := (values[u] + values[u+1]) / 2
	fmt.Printf("%10s:%5d\n", "Lower", lower)
	fmt.Printf("%10s:%5d\n", "Median", median)
	fmt.Printf("%10s:%5d\n", "Upper", upper)

	sd := float64(0)
	for i := 0; i < len(values); i++ {
		sd += math.Pow(math.Abs(float64(values[i])-avg), 2)
	}
	sd = math.Sqrt(sd / float64(len(values)))
	fmt.Printf("%10s:%5.2f\n", "SD", sd)
	fmt.Printf("%10s:%5.2f\n", "SD", sd)
	ts := time.Duration(ts2 - ts1)
	fmt.Printf("%10s:%v\n", "elapse", ts.String())
}

func SumInt(values []int32) int32 {
	var ttl int32
	for _, v := range values {
		ttl += v
	}
	return ttl
}
