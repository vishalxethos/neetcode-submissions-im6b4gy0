func productExceptSelf(nums []int) []int {

	output := make([]int, 0)

	n := len(nums)

	output = append(output, 1)


	for i := 1; i <= n-1; i++ {
		output = append(output, output[i-1]*nums[i-1])
	}


	right := 1

	for j := n - 2; j >= 0; j-- {

		right = nums[j+1] * right
		output[j] = output[j] * right

	}

	return output
}
