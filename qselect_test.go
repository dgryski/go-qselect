package qselect

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"testing/quick"
)

func TestSelect(t *testing.T) {

	f := func(a []int) bool {

		if len(a) == 0 {
			return true
		}

		b := make([]int, len(a))
		copy(b, a)
		sort.Ints(b)

		pos := rand.Intn(len(a))

		Select(sort.IntSlice(a), pos)

		if a[pos] != b[pos] {
			return false
		}

		for i := 0; i < pos; i++ {
			if a[i] > a[pos] {
				return false
			}
		}

		for i := pos; i < len(a); i++ {
			if a[i] < a[pos] {
				return false
			}
		}

		return true
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func benchmarkSort(b *testing.B, data []int) {
	arr := make([]int, len(data))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(arr, data)
		sort.Ints(arr)
	}
}

func benchmarkSelect(b *testing.B, data []int) {
	arr := make([]int, len(data))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(arr, data)
		Select(sort.IntSlice(arr), len(arr)/2)
	}
}

func BenchmarkSort(b *testing.B) {
	runBenchmarks(b, benchmarkSort)
}

func BenchmarkSelect(b *testing.B) {
	runBenchmarks(b, benchmarkSelect)
}

func runBenchmarks(b *testing.B, bench func(b *testing.B, data []int)) {

	var sizes = []int{
		10,
		20,
		50,
		100,
		1024,
		2048,
		4096,
		8192,
		10000,
		20000,
		1e5,
		2e5,
		1e6,
	}

	for _, sz := range sizes {
		rand.Seed(0)
		data := make([]int, sz)
		for i := range data {
			data[i] = rand.Int()
		}

		b.Run(strconv.Itoa(sz), func(b *testing.B) { bench(b, data) })
	}
}
