package leetcode

// Given a string, find the length of the longest substring without repeating characters.
//
// Examples:
// Given "abcabcbb", the answer is "abc", which the length is 3.
// Given "bbbbb", the answer is "b", with the length of 1.
// Given "pwwkew", the answer is "wke", with the length of 3.
// Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
//
func lengthOfLongestSubstring(s string) int {
	hashmap := map[rune]int{}
	max := 0
	left := 0
	for right, r := range s {
		if prev, ok := hashmap[r]; ok && prev >= left {
			left = prev + 1
		}
		hashmap[r] = right
		if right-left+1 > max {
			max = right - left + 1
		}
	}
	return max
}