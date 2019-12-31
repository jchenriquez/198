package main

import "fmt"

func sumBySkip (nums []int, i, j int, seen map[int]int) int {
	if i >= j {
		return 0
	}

	if value, saw := seen[i]; saw {
		return value
	}

	return nums[i] + sumBySkip(nums, i+2, j, seen)
}

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	highestBounty := 0
	seen := make(map[int]int)

	for i := len(nums)-1; i >= 0; i-- {
		var rightSum int

		if i <= len(nums) - 2 {
			for j := i+2; j < len(nums); j++ {
				tSum := sumBySkip(nums, j, len(nums), seen)
				if rightSum < tSum {
					rightSum = tSum
				}
			}
		}

		if rightSum+nums[i] > highestBounty {
			highestBounty = rightSum + nums[i]
		}

		seen[i] = rightSum + nums[i]
	}

	return highestBounty
}


func main() {
	fmt.Printf("highest bounty is %v\n", rob([]int{6,3,10,8,2,10,3,5,10,5,3}))
}
