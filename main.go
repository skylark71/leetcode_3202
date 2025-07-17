package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	k := 2
	fmt.Println(maximumLength(arr, k))
}

func maximumLength(nums []int, k int) int {
	n := len(nums)
	dp := make([]map[int]int, n)

	maxLen := 1

	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			mod := (nums[j] + nums[i]) % k
			prevLen := dp[j][mod]
			if prevLen == 0 {
				prevLen = 1 // стартовая длина
			}
			dp[i][mod] = max(dp[i][mod], prevLen+1)
			maxLen = max(maxLen, dp[i][mod])
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
