
func productExceptSelf(nums []int) []int {

	output := make([]int, len(nums))

	for index := range nums {
		if index == 0 {
			output[0] = 1
			continue
		}

		output[index] = nums[index-1] * output[index-1]

	}

	runningTotal := 1

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			continue
		}
		runningTotal =  runningTotal * nums[i+1]
		output[i] = output[i] * runningTotal
		

	}

	return output

}
