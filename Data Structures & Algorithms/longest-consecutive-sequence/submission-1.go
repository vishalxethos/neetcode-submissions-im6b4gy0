func longestConsecutive(nums []int) int {

	hashMap := make(map[int]bool)
	maxSequence := 1

	if len(nums) == 0 {

		return 0
	}


	for _,value := range nums {

	       hashMap[value] = true
	}

	for _, value := range nums {
			currentSequence := 1
			if _,exists := hashMap[value-1] ; !exists {
					currentValue := value

				
					for  hashMap[currentValue+1]  {
						currentSequence += 1
						currentValue += 1
					}

					maxSequence = max(maxSequence,currentSequence)
					

			} 
	}

 return maxSequence
}
