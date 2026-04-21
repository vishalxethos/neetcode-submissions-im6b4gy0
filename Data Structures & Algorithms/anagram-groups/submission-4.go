func groupAnagrams(strs []string) [][]string {

	mapper := make(map[[26]byte][]string)

	result := [][]string{}

	for _, value := range strs {
		key := [26]byte{}

		for i := 0; i < len(value); i++ {
			key[value[i]-97] += 1
		}

		mapper[key] = append(mapper[key], value)

	}

	for _,value := range mapper {

		result = append(result, value)

	}

	return result

}
