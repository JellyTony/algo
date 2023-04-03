package longest_substring_without_repeating_characters

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	substring := lengthOfLongestSubstring("abcabcbb")
	t.Log(substring)
}
