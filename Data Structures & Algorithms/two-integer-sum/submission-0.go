func twoSum(nums []int, target int) []int {
	results := make(map[int]int)

	for i, value := range nums {
		j, ok := results[target-value]

		if ok {
			return []int{j, i}
		}

		results[value] = i
	}

	return []int{-1, -1}
}

