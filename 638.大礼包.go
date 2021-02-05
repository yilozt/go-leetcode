/*
 * @lc app=leetcode.cn id=638 lang=golang
 *
 * [638] 大礼包
 */

// @lc code=start
package main

import "fmt"

var (
	cache map[[6]int]int
)

// 返回两者间的最小值
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	// 初始化缓存
	cache = make(map[[6]int]int)
	var fixSizeNeeds [6]int
	for i := 0; i < len(needs); i++ {
		fixSizeNeeds[i] = needs[i]
	}
	fmt.Println(fixSizeNeeds)
	return helper(price, special, fixSizeNeeds)
}

func helper(price []int, special [][]int, needs [6]int) int {
	if result, ok := cache[needs]; ok {
		return result
	}
	// 以不买礼包的价格作为基准
	result := calculateTotalMoney(price, needs)
	// 遍历所有礼包
	for _, gift := range special {
		// 拷贝一份原来的代购清单
		cloneNeeds := needs
		// 标记是否买超
		flag := true
		// 检查买这个礼包后是否会超出代购清单
		for index := range price {
			// 买了礼包后剩余要买的物品数量
			needNow := cloneNeeds[index] - gift[index]
			// 如果买超了就不买了
			if needNow < 0 {
				flag = false
				break
			}
			// 否则保存新的需购买物品数量
			cloneNeeds[index] = needNow
		}
		if flag {
			result = min(result, gift[len(gift)-1]+helper(price, special, cloneNeeds))
		}
	}
	cache[needs] = result
	return result
}

// 通过价格和商品得到总价格
func calculateTotalMoney(price []int, needs [6]int) int {
	result := 0
	for i := range price {
		result += price[i] * needs[i]
	}
	return result
}

// @lc code=end
