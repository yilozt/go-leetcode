/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] 验证IP地址
 */

// @lc code=start
package main

import (
	"strconv"
	"strings"
)

const (
	IPV4 = "IPv4"
	IPV6 = "IPv6"
	Err  = "Neither"
)

func checkByte(b byte) bool {
	if (b >= '0' && b <= '9') || (b >= 'A' && b <= 'F') || (b >= 'a' && b <= 'f') {
		return true
	}
	return false
}

func validIPAddress(IP string) string {
	// 尝试判断IPV4
	if strings.Contains(IP, ".") {
		strs := strings.Split(IP, ".")
		if len(strs) != 4 {
			return Err
		}
		for _, str := range strs {
			if len(str) > 1 && strings.HasPrefix(str, "0") {
				return Err
			}
			strInt, err := strconv.Atoi(str)
			if err != nil || strInt > 255 || strInt < 0 {
				return Err
			}
		}
		return IPV4
	} else {
		strs := strings.Split(IP, ":")
		if len(strs) != 8 {
			return Err
		}
		for _, str := range strs {
			if len(str) > 4 || len(str) == 0 {
				return Err
			}
			for i := 0; i < len(str); i++ {
				if !checkByte(str[i]) {
					return Err
				}
			}
		}
		return IPV6
	}
}

// @lc code=end
