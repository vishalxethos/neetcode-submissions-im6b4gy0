func twoSum(nums []int, target int) []int {

	mapper := make(map[int]int)

	for index := range nums {

			if _, exists := mapper[target - nums[index] ]; exists {
				reverseIndex := mapper[target - nums[index]]
				return []int{reverseIndex,index}
			}

			mapper[nums[index]] = index
			fmt.Println(mapper)
	}

	return []int{}
    
}
