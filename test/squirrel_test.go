package test

import (
	. "github.com/edkarlsson/go-squirrelnoise/pkg/squirrelnoise"
	"strconv"
	"testing"
)

// Can probably make this unique test better and expand it for the other functions
func TestSquirrelNoise5(t *testing.T) {
	var want uint32 = 377036288
	if got := SquirrelNoise5(0, 0); got != want {
		t.Errorf("SquirrelNoise5(0,0) = %d; want %d", got, want)
	}
	arrLen := 100000
	uniqueMap := make(map[string]uint)
	var res uint32 = 0
	for i := 0; i < arrLen; i++ {
		res = SquirrelNoise5(int32(i), res)
		randK := strconv.FormatInt(int64(res), 10)

		if val, ok := uniqueMap[randK]; ok {
			uniqueMap[randK]++
		} else if val == 1 {
			t.Errorf("SquirrelNoise5(%d, 0) = %s was not unique", i, randK)
		} else {
			uniqueMap[randK] = 1
		}
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
