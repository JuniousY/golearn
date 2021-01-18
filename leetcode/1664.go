package main

import (
	"fmt"
)

// 196 ms faster than 16.15%
func waysToMakeFair(nums []int) int {
	last := len(nums) - 1
	evenSum := make([]int, len(nums))
	oddSum := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			evenSum[0] = v
			continue
		}
		if i == 1 {
			evenSum[1] = evenSum[0]
			oddSum[1] = v
			continue
		}
		if isEven(i) {
			evenSum[i] = evenSum[i-2] + v
			oddSum[i] = oddSum[i-1]
		} else {
			oddSum[i] = oddSum[i-2] + v
			evenSum[i] = evenSum[i-1]
		}
	}
	count := 0
	for i, v := range nums {
		if i == 0 {
			if evenSum[last]-v == oddSum[last] {
				count++
			}
			continue
		}
		if i == 1 {
			if evenSum[last] - nums[0] == oddSum[last]-v + nums[0]{
				count++
			}
			continue
		}
		if isEven(i) {
			if evenSum[i-2]+(oddSum[last]-oddSum[i-1]) ==
				oddSum[i-1]+(evenSum[last]-evenSum[i]) {
				count++
			}
		} else {
			if evenSum[i-1]+(oddSum[last]-oddSum[i]) ==
				oddSum[i-2]+(evenSum[last]-evenSum[i-1]) {
				count++
			}
		}
	}
	return count
}

func isEven(n int) bool {
	return n&1 == 0
}

// todo 更好的方法

func main() {
	//nums := []int{2, 1, 6, 4}
	//nums := []int{1, 1, 1}
	nums := []int{1, 2, 3}
	fmt.Println(waysToMakeFair(nums))
}
