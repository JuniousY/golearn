package leetcode

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
			if evenSum[last]-nums[0] == oddSum[last]-v+nums[0] {
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

// copy from leetcode 116 ms
func waysToMakeFair2(nums []int) int {
	length := len(nums)
	odd_sum := 0
	even_sum := 0
	result := 0
	var odd_sum_list []int = make([]int, length)
	var even_sum_list []int = make([]int, length)
	for index, num := range nums {
		if index%2 == 1 {
			odd_sum += num
		} else {
			even_sum += num
		}
		odd_sum_list[index] = odd_sum
		even_sum_list[index] = even_sum
	}

	for index, _ := range nums {
		current_back_odd_sum := (even_sum - even_sum_list[index])
		current_back_even_sum := (odd_sum - odd_sum_list[index])

		if index > 0 {
			if current_back_odd_sum+odd_sum_list[index-1] == current_back_even_sum+even_sum_list[index-1] {
				result += 1
			}

		} else {
			if current_back_odd_sum == current_back_even_sum {
				result += 1
			}
		}
	}
	return result
}

func main() {
	//nums := []int{2, 1, 6, 4}
	//nums := []int{1, 1, 1}
	nums := []int{1, 2, 3}
	fmt.Println(waysToMakeFair(nums))
}
