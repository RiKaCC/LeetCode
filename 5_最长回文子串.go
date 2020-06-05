// 思路：从第二个元素开始，向两边扩散，寻找以他为中心的回文串。

func LongestPalindrome(s string) string {
	ln := len(s)
	if ln == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < ln; i++ {
		s1, e1 := findPalindrome(s, i, i)   // s长度为奇数
		s2, e2 := findPalindrome(s, i, i+1) // s长度为偶数
		if e1-s1 > end-start {
			start, end = s1, e1
		}
		if e2-s2 > end-start {
			start, end = s2, e2
		}
	}
	return s[start : end+1]
}

func findPalindrome(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}
	return left + 1, right - 1
}
