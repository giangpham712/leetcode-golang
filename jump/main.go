package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 1, 4}
	result := jump(nums)

	fmt.Println(result)
}

func jump2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	max := 0
	current := 0
	count := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > max {
			max = i + nums[i]
		}
		if current == i {
			current = max
			count++
		}
		if current > len(nums)-1 {
			return count
		}
	}
	return count
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	minMap := make(map[int]int)
	minMap[len(nums)-1] = 0

	result := minJump(nums, 0, minMap)

	return result
}

func minJump(nums []int, from int, minMap map[int]int) int {
	min, exists := minMap[from]
	if exists {
		return min
	} else {
		min = len(nums) - from
	}

	for i := nums[from]; i > 0; i-- {
		if from+i > len(nums)-1 {
			continue
		}

		minI := minJump(nums, from+i, minMap)
		if minI < min {
			min = minI
		}
	}

	minMap[from] = min + 1
	return min + 1
}
