package main

import (
	"fmt"
	"time"
)

var z int

func main() {
	//numStock := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	numStock := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		numStock[i] = i
	}
	nums := []int{2, 4, 5, 7, 10, 12, 14, 15, 16, 20, 31, 32, 34, 45, 48, 49, 50, 68, 79, 82, 91, 93, 99, 498, 998, 999998}

	result := make([]int, len(numStock)-len(nums))

	now := time.Now()
	k := 0
	for i := 0; i < len(numStock); i++ {
		for j := 0; j < len(nums); j++ {
			if numStock[i] == nums[j] {
				break
			}
			if j == len(nums)-1 {
				//fmt.Printf("%d %d %d %d %d\r\n", i, j, k, numStock[i], nums[j])
				result[k] = numStock[i]
				k++
			}
		}
	}

	fmt.Printf("took %v to get %v\r\n", time.Since(now), result[(len(result)-10):])

	now = time.Now()
	result = findDiffOnTwoList(numStock, nums)
	fmt.Printf("took %v to get %v\r\n", time.Since(now), result[(len(result)-10):])

}

func findDiffOnTwoList(master []int, nums []int) []int {
	masterLen := len(master)
	if masterLen == 1 {
		isDiff := true
		for _, n := range nums {
			if n == master[0] {
				isDiff = false
				break
			}
		}
		if isDiff {
			return []int{master[0]}
		}
		return []int{}
	}

	left := findDiffOnTwoList(master[:masterLen/2], nums)
	right := findDiffOnTwoList(master[masterLen/2:], nums)

	for _, r := range right {
		left = append(left, r)
	}
	return left
}

func sumElementRecursion() {
	nums := []int{6, 4, 2, 8, 1, 14, 5, 7, 9, 3, 4, 5, 1, 24, 5, 3, 6, 2, 9, 1, 13}
	now := time.Now()
	var result int
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	fmt.Printf("took %v to get %d \r\n", time.Since(now), result)
	now = time.Now()
	i := 0
	result = sumElement(nums, i)
	fmt.Printf("took %v to get %d \r\n", time.Since(now), result)
}

func sumElement(numbers []int, i int) int {
	if i == len(numbers) {
		return 0
	}

	return sumElement(numbers, i+1) + numbers[i]
}
