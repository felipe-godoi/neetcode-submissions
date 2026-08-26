func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	occurrences := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		occurrences[s[i]]++
		occurrences[t[i]]--
	}

	for _, value := range occurrences {
		if value != 0 {
			return false
		}
	}

	return true
}

