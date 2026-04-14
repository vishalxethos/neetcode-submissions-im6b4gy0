func groupAnagrams(strs []string) [][]string {

	mapper := make(map[[26]int][]string)

	result := [][]string{}

	for _, value := range strs {
		key := [26]int{}

		for i := 0; i < len(value); i++ {
			key[value[i]-97] += 1
		}

		mapper[key] = append(mapper[key], value)

	}

	for value := range mapper {

		result = append(result, mapper[value])

	}

	return result

}
