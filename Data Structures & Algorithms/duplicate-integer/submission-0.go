func hasDuplicate(nums []int) bool {
    	mapper := make(map[int]struct{})

	for index := range nums {

		if _, exists := mapper[nums[index]]; exists {
			return true
		}

		mapper[nums[index]] = struct{}{}
	}

	return false
}
