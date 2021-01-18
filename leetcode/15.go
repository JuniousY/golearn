package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	first := 0
	var result [][]int
	for first < len(nums) && nums[first] <= 0 {
		for first > 0 && first < len(nums) && nums[first-1] == nums[first] {
			first++
		}
		if first >= len(nums) {
			continue
		}
		firstVal := nums[first]
		second := first + 1
		third := len(nums) - 1
		for second < third {
			if sum := firstVal + nums[second] + nums[third]; sum == 0 {
				result = append(result, []int{firstVal, nums[second], nums[third]})
				for second < third-1 && nums[second+1] == nums[second] {
					second++
				}
				for third > second+1 && nums[third-1] == nums[third] {
					third--
				}
				second++
				third--
			} else if sum < 0 {
				for second < third-1 && nums[second+1] == nums[second] {
					second++
				}
				second++
			} else {
				for third > second+1 && nums[third-1] == nums[third] {
					third--
				}
				third--
			}
		}
		first++
	}
	return result
}

// copy from leetcode 20ms
func threeSum2(nums []int) [][]int {
	// After sorting, you can search by law
	sort.Ints(nums)

	res := [][]int{}
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		// Avoid adding duplicate results
		// i>0 is to prevent nums[i-1] from overflowing
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				// the smaller l needs to be larger
				l++
			case s > 0:
				// The larger r needs to be smaller
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// To avoid repeated additions, both l and r need to be moved to different elements.
				l, r = next(nums, l, r)
			}
		}
	}
	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}
	return l, r
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//var nums []int
	//nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))
}

