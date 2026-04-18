func topKFrequent(nums []int, k int) []int {
	bucket := make([][]int, len(nums)+1)

	counter := k

	result := make([]int, 0, k)

	freqMap := make(map[int]int, 0)

	for _, elem := range nums {

		freqMap[elem]++
	}

//	fmt.Println(freqMap)

	for key, value := range freqMap {

		bucket[value] = append(bucket[value], key)
	}
//	fmt.Println(bucket)

	for i := len(bucket) - 1; i > 0; i-- {

		if counter <= 0 {
			break
		}
		if (len(bucket[i])) > 0 {
			result = append(result, bucket[i]...)
			counter -= len(bucket[i])
		}

	}

	return result[0:k]

}
