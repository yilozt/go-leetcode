/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

package main

import "fmt"

// @lc code=start
func longestPalindrome(s string) int {
	// 统计字母出现的次数
	count := make(map[rune]int, 0)
	for _, c := range s {
		count[c]++
	}

	max_odd_count := 0
	ans := 0
	for _, c := range count {
		ans += c / 2 * 2
		if c%2 == 1 && ans%2 == 0 {
			ans++
		}
	}
	ans += max_odd_count

	return ans
}

// @lc code=end

func main() {
	fmt.Println(longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}
