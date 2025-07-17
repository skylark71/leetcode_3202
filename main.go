package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	k := 2
	fmt.Println(maximumLength(arr, k))
}

func maximumLength(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n*k)
	maxLen := 1

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			mod := (nums[j] + nums[i]) % k
			prev := dp[j*k+mod]
			if prev == 0 {
				prev = 1
			}
			idx := i*k + mod
			if prev+1 > dp[idx] {
				dp[idx] = prev + 1
				if dp[idx] > maxLen {
					maxLen = dp[idx]
				}
			}
		}
	}
	return maxLen
}
