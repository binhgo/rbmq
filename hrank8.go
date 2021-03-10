package main

func maxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxSum, res, p := nums[0], 0, 0

	for p < len(nums) {

		res += nums[p]

		if res > maxSum {
			maxSum = res
		}

		if res < 0 {
			res = 0
		}

		p++

	}
	return maxSum
}

func main() {

}
