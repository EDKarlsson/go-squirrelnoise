package test

import (
	"fmt"
	. "github.com/edkarlsson/go-squirrelnoise"
	"strconv"
	"testing"
)

// Can probably make this unique test better and expand it for the other functions
func TestSquirrelNoise5(t *testing.T) {
	var want uint32 = 377036288
	if got := SquirrelNoise5(0, 0); got != want {
		t.Errorf("SquirrelNoise5(0,0) = %d; want %d", got, want)
	}
	arrLen := 1000000
	uniqueMap := make(map[string]uint)
	nonUnique := 0
	var res uint32 = 0
	for i := 0; i < arrLen; i++ {
		// prevRes := res
		res = SquirrelNoise5(int32(i), res)
		randK := strconv.FormatInt(int64(res), 10)

		if _, ok := uniqueMap[randK]; ok {
			uniqueMap[randK]++
			if uniqueMap[randK] > 1 {
				// fmt.Printf("Found non unique SquirrelNoise5(%d, %d)=%s, count=%d\n", i, prevRes, randK, uniqueMap[randK])
				nonUnique += 1
			}
		} else {
			uniqueMap[randK] = 1
		}
	}
	fmt.Printf("Number of random values generated = %d\n", arrLen)
	fmt.Printf("Number of repeated values found = %d\n", nonUnique)
	fmt.Printf("Percent of collisions = %1.6f\n", float64(nonUnique)/float64(arrLen))
	if float64(nonUnique)/float64(arrLen) > 1 {
		t.Errorf("More than 1%% of the values of %d were repeated", arrLen)
	}
}

func BenchmarkSquirrelNoise5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquirrelNoise5(int32(i), 0)
	}
}
func BenchmarkGet1dNoiseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get1dNoiseUint(int32(i), 0)
	}
}
func BenchmarkGet2dNoiseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get2dNoiseUint(int32(i), int32(i), 0)
	}
}
func BenchmarkGet3dNoiseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get3dNoiseUint(int32(i), int32(i), int32(i), 0)
	}
}
func BenchmarkGet4dNoiseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get4dNoiseUint(int32(i), int32(i), int32(i), int32(i), 0)
	}
}
func BenchmarkGet1dNoiseZeroToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get1dNoiseZeroToOne(int32(i), 0)
	}
}
func BenchmarkGet2dNoiseZeroToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get2dNoiseZeroToOne(int32(i), int32(i), 0)
	}
}
func BenchmarkGet3dNoiseZeroToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get3dNoiseZeroToOne(int32(i), int32(i), int32(i), 0)
	}
}
func BenchmarkGet4dNoiseZeroToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get4dNoiseZeroToOne(int32(i), int32(i), int32(i), int32(i), 0)
	}
}
func BenchmarkGet1dNoiseNegToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get1dNoiseNegOneToOne(int32(i), 0)
	}
}
func BenchmarkGet2dNoiseNegToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get2dNoiseNegOneToOne(int32(i), int32(i), 0)
	}
}
func BenchmarkGet3dNoiseNegToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get3dNoiseNegOneToOne(int32(i), int32(i), int32(i), 0)
	}
}
func BenchmarkGet4dNoiseNegToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get4dNoiseNegOneToOne(int32(i), int32(i), int32(i), int32(i), 0)
	}
}
