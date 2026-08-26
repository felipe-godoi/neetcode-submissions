func hasDuplicate(nums []int) bool {
    occurrences := make(map[int]struct{})

	for _, value := range nums {
		_, ok := occurrences[value]

		if !ok {
			occurrences[value] = struct{}{}
		} else {
			return true
		}
	}

	 return false
}
