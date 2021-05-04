package test

import (
	"testing"
)

func BenchmarkConditionalIf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pronounciationIf(n)
	}
}

func BenchmarkConditionalIfElse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pronounciationIfElse(n)
	}
}

func BenchmarkConditionalSwitch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pronounciationSwitch(n)
	}
}

//func BenchmarkConditional(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		pronounciationSwitch(n)
//	}
//}

func pronounciationIf(num int) string {
	var result string
	if num == 2 {
		result = "dua"
	}
	if num == 3 {
		result = "tiga"
	}
	if num == 4 {
		result = "empat"
	}
	if num == 5 {
		result = "lima"
	}
	if num == 6 {
		result = "enam"
	}
	if num == 7 {
		result = "tujuh"
	}
	if num == 8 {
		result = "delapan"
	}
	if num == 9 {
		result = "sembilan"
	}
	if num == 10 {
		result = "sepuluh"
	}
	if num == 1 {
		result = "satu"
	}

	return result
}

func pronounciationSwitch(num int) string {
	var result string
	switch num {
	case 2:
		result = "dua"
	case 3:
		result = "tiga"
	case 4:
		result = "empat"
	case 5:
		result = "lima"
	case 6:
		result = "enam"
	case 7:
		result = "tujuh"
	case 8:
		result = "delapan"
	case 9:
		result = "sembilan"
	case 10:
		result = "sepuluh"
	case 1:
		result = "satu"
	}

	return result
}

func pronounciationIfElse(num int) string {
	var result string
	if num == 2 {
		result = "dua"
	} else if num == 3 {
		result = "tiga"
	} else if num == 4 {
		result = "empat"
	} else if num == 5 {
		result = "lima"
	} else if num == 6 {
		result = "enam"
	} else if num == 7 {
		result = "tujuh"
	} else if num == 8 {
		result = "delapan"
	} else if num == 9 {
		result = "sembilan"
	} else if num == 10 {
		result = "sepuluh"
	} else if num == 1 {
		result = "satu"
	}

	return result
}
