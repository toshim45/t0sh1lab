package main

import (
	"fmt"
	"time"
)

var z int

func main() {
	//	nums := []int{6, 4, 2, 8, 1, 14, 5, 7, 9, 67, 3, 11, 15, 26, 36, 34, 13, 19, 69}
	nums := []int{6, 4, 2, 8, 1, 14, 5, 7, 9, 67, 3, 11, 15, 26, 36, 34, 13, 19, 69, 10, 23, 35, 45, 78, 89, 90, 98, 76, 65, 54, 43, 32, 21, 106, 104, 102, 108, 101, 114, 105, 107, 109, 167, 103, 111, 115, 126, 136, 134, 113, 119, 169, 110, 123, 135, 145, 178, 189, 190, 198, 176, 165, 154, 143, 132, 221, 206, 204, 202, 208, 201, 214, 205, 207, 209, 267, 203, 211, 215, 226, 236, 234, 213, 219, 269, 210, 223, 235, 245, 278, 289, 290, 298, 276, 265, 254, 243, 232, 321, 306, 304, 302, 308, 301, 314, 305, 307, 309, 367, 303, 311, 315, 326, 336, 334, 313, 319, 369, 310, 323, 335, 345, 378, 389, 390, 398, 376, 365, 354, 343, 332, 321}
	numsLen := len(nums)
	now := time.Now()
	sortedNums := mergeSort(nums)
	PrintIntegers(sortedNums)
	fmt.Println("iteration cost : ", z, " length: ", numsLen, " took: ", time.Since(now))
	z = 0
	now = time.Now()
	sortedNums = bubleSort(nums)
	PrintIntegers(sortedNums)
	fmt.Println("iteration cost : ", z, " length: ", numsLen, " took: ", time.Since(now))
}

func bubleSort(numbers []int) []int {
	numLen := len(numbers)

	for i := 0; i < numLen; i++ {
		for j := 0; j < numLen; j++ {
			z++
			if numbers[i] < numbers[j] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}

	return numbers
}

func mergeSort(numbers []int) []int {
	z++
	numLen := len(numbers)
	//	fmt.Println("processing len > ", numLen)
	//	PrintIntegers(numbers)

	if numLen < 2 {
		return numbers
	}

	halfNumLen := numLen / 2

	left := mergeSort(numbers[0:halfNumLen])
	right := mergeSort(numbers[halfNumLen:numLen])

	i, j := 0, 0
	result := make([]int, numLen)

	//	fmt.Println("left > ")
	//	PrintIntegers(left)
	//	fmt.Println("right > ")
	//	PrintIntegers(right)

	//merge
	for k := 0; k < numLen; k++ {
		//		fmt.Printf("%d %d %d < %d < %v\r\n", k, i, j, numLen, result)
		if i >= len(left) && j < len(right) {
			result[k] = right[j]
			j++
		} else if i < len(left) && j >= len(right) {
			result[k] = left[i]
			i++
		} else if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}

	return result
}
