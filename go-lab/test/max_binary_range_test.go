package test

import (
	"strconv"
	"testing"
)

var cases = []struct {
	Number        int64
	ExpectedRange int
}{
	{145, 3},
	{3332110, 7},
	{2147483647, 0},
	{2000480007, 5},
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

	// fmt.Printf("idxs-m: %v\n", oneIdxs)
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

// func divideMaxBinaryRange(in int64) (out int) {
// 	return
// }

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
}
