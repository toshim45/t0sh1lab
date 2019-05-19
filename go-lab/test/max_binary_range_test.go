package test

// go test -bench="BenchmarkMaxBinaryRange*" -benchmem
// goos: darwin
// goarch: amd64
// BenchmarkMaxBinaryRangeNormal-4          	 2000000	       685 ns/op	     968 B/op	       9 allocs/op
// BenchmarkMaxBinaryRangeMerge-4           	 1000000	      1264 ns/op	    1472 B/op	      26 allocs/op
// BenchmarkMaxBinaryRangeDivideConquer-4   	 1000000	      1826 ns/op	    1552 B/op	      35 allocs/op

import (
	"strconv"
	"testing"
)

var cases = []struct {
	Number        int64
	ExpectedRange int
}{
	{145, 3},
	{1665, 6},
	{3332110, 7},
	{2147483647, 0},
	{2000480007, 5},
}

func calculateMaxRange(oneIdxs []int) (out int) {
	rangeIdxs := make([]int, len(oneIdxs)-1)
	for i, idx := range oneIdxs {
		if i == 0 {
			continue
		}

		rangeIdxs = append(rangeIdxs, idx-oneIdxs[i-1]-1)
	}

	for _, rangeIdx := range rangeIdxs {
		if rangeIdx > out {
			out = rangeIdx
		}
	}
	return
}

func normalMaxBinaryRange(in int64) (out int) {
	bf := strconv.FormatInt(in, 2)
	bs := []rune(bf)
	oneIdxs := make([]int, 0)
	for i, f := range bs {
		if f == '1' {
			oneIdxs = append(oneIdxs, i)
		}
	}

	// fmt.Printf("idxs-n: %v\n", oneIdxs)
	return calculateMaxRange(oneIdxs)
}

func mergeMaxBinaryRange(in int64) (out int) {
	bf := strconv.FormatInt(in, 2)
	bs := []rune(bf)
	oneIdxsLeft := make([]int, 0)
	oneIdxsRight := make([]int, 0)
	lbs := len(bs) - 1

	for i := 0; i <= lbs; i++ {
		if bs[i] == '1' {
			oneIdxsLeft = append(oneIdxsLeft, i)
		}
		if bs[lbs] == '1' {
			oneIdxsRight = append([]int{lbs}, oneIdxsRight...)
		}
		lbs--
		// fmt.Printf("idxs-m: %v %v\n", oneIdxsLeft, oneIdxsRight)
	}

	oneIdxs := append(oneIdxsLeft, oneIdxsRight...)

	return calculateMaxRange(oneIdxs)
}

func divideMaxBinaryRange(in int64) (out int) {
	bf := strconv.FormatInt(in, 2)
	bs := []rune(bf)
	oneIdxs := findIdxOne(bs, 0)
	// fmt.Printf("idxs-n: %v\n", oneIdxs)
	return calculateMaxRange(oneIdxs)
}

func findIdxOne(in []rune, startIdx int) (out []int) {
	// fmt.Printf("raw-divide: %v %d\n", in, startIdx)
	l := len(in)
	if l == 2 {
		if in[0] == '1' {
			out = append(out, startIdx)
		}
		if in[1] == '1' {
			out = append(out, startIdx+1)
		}
		return
	} else if l == 1 {
		if in[0] == '1' {
			out = append(out, startIdx)
		}
		return
	}

	h := l / 2
	out = append(findIdxOne(in[0:h], startIdx), out...)
	out = append(out, findIdxOne(in[h:l], startIdx+h)...)
	return
}

func TestMaxBinaryRange(t *testing.T) {
	for _, c := range cases {
		mbr := normalMaxBinaryRange(c.Number)
		if mbr != c.ExpectedRange {
			t.Errorf("got range %d should be %d", mbr, c.ExpectedRange)
		}
	}
	for _, c := range cases {
		mbr := mergeMaxBinaryRange(c.Number)
		if mbr != c.ExpectedRange {
			t.Errorf("got range %d should be %d", mbr, c.ExpectedRange)
		}
	}
	for _, c := range cases {
		mbr := divideMaxBinaryRange(c.Number)
		if mbr != c.ExpectedRange {
			t.Errorf("got range %d should be %d", mbr, c.ExpectedRange)
		}
	}
}

func BenchmarkMaxBinaryRangeNormal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		normalMaxBinaryRange(2000480007)
	}
}
func BenchmarkMaxBinaryRangeMerge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mergeMaxBinaryRange(2000480007)
	}
}
func BenchmarkMaxBinaryRangeDivideConquer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		divideMaxBinaryRange(2000480007)
	}
}
