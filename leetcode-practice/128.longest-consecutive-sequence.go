// https://leetcode.com/problems/longest-consecutive-sequence/




func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}

	maxLen := 0
	for k, _ := range m {
		if _, ok := m[k-1]; ok {
			continue
		}
		current := k
		length := 1
		for {
			if _, ok := m[current+1]; !ok {
				break
			}
			length++
			current++
		}
		if length > maxLen {
			maxLen = length
		}

	}
	return maxLen

}
