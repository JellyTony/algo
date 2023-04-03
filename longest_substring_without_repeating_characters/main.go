package longest_substring_without_repeating_characters

// lengthOfLongestSubstring 获取最长无重复子串的长度
func lengthOfLongestSubstring(s string) int {
	n := len(s)                                   // 字符串长度
	result := 0                                   // 最长无重复子串的长度
	charIndex := make([]int, 256)                 // 用数组记录每个字符最近一次出现的位置
	for left, right := 0, 0; right < n; right++ { // 左右指针法遍历字符串
		c := s[right]                      // 当前字符
		left = max(charIndex[c], left)     // 更新左指针的位置
		result = max(result, right-left+1) // 更新最长无重复子串的长度
		charIndex[c] = right + 1           // 更新当前字符最近一次出现的位置
	}

	return result // 返回最长无重复子串的长度
}

func max(x, y int) int {
	if x > y { // 返回两个数中的较大值
		return x
	}
	return y
}
