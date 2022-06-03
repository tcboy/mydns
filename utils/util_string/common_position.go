package util_string

import (
	"unicode/utf16"
)

// 动态规划算公共子序列边界
func StrCommentPositions(target, str string) [][]int {
	positions := make([][]int, 0)

	targetRune := toUtf16(target)
	strRune := toUtf16(str)

	targetLen := len(targetRune) + 1
	strLen := len(strRune) + 1

	dp := make([][]int, targetLen, targetLen)
	for i, _ := range dp {
		dp[i] = make([]int, strLen, strLen)
	}

	for i := 1; i < targetLen; i++ {
		for j := 1; j < strLen; j++ {
			if targetRune[i-1] == strRune[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	lcsRune := make([]uint16, 0)
	for i, j := targetLen-1, strLen-1; i > 0 && j > 0; {
		if targetRune[i-1] == strRune[j-1] && dp[i][j] == dp[i-1][j-1]+1 {
			lcsRune = append(lcsRune, targetRune[i-1])
			i -= 1
			j -= 1
		} else if targetRune[i-1] != strRune[j-1] && dp[i-1][j] > dp[i][j-1] {
			i -= 1
		} else {
			j -= 1
		}
	}
	reverse(lcsRune)

	lcsPos, left, right := 0, -1, -1
	for i := 0; i < targetLen-1; i++ {
		if lcsPos < len(lcsRune) && i == 0 && targetRune[i] == lcsRune[lcsPos] {
			left, right = 0, 0
			lcsPos += 1
		} else if lcsPos < len(lcsRune) && lcsPos == 0 && targetRune[i] == lcsRune[lcsPos] {
			left, right = i, i
			lcsPos += 1
		} else if lcsPos < len(lcsRune) && targetRune[i] == lcsRune[lcsPos] {
			if targetRune[i-1] == lcsRune[lcsPos-1] && left >= 0 && right >= 0 {
				right += 1
			} else {
				left, right = i, i
			}
			lcsPos += 1
		} else if left >= 0 && right >= 0 {
			positions = append(positions, []int{left, right})
			left, right = -1, -1
		}
	}
	if left >= 0 && right >= 0 {
		positions = append(positions, []int{left, right})
	}

	return positions
}

func toUtf16(str string) []uint16 {
	return utf16.Encode([]rune(str))
}

func reverse(s []uint16) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
