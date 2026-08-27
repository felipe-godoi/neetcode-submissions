func groupAnagrams(strs []string) [][]string {
	frequencies := make(map[[26]int][]string)

	for _, str := range strs {
		var freqArray [26]int

		for _, c := range str {
			freqArray[c-97]++
		}

		frequencies[freqArray] = append(frequencies[freqArray], str)
	}

	values := make([][]string, 0, len(frequencies))

	for _, value := range frequencies {
		values = append(values, value)
	}

	return values
}
