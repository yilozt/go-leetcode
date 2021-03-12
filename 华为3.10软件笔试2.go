/*
题目描述：
司机小王到仓库拉货。仓库管理员已经提前将货物打包装箱,并摆成一排。包装箱从1
始编号。各个包装箱存放的货物数组成一个集合M={M1, M2. ..., Mn}。
货车一次最多运送K件货物。为了最大化动力,小王想一次从中挑选出K的整数倍件货物,
再分批运输。仓库管理员为了方便管理,要求小王必须选择编号连续的包装箱,如可选择
1、2、3号箱,不能选择2、4、6号箱。
如果运输K整数倍件货物,请帮小王计算有多少种挑选方式。

输入描述：
包装箱数为N,货车一次运送货物的最大数为K
各箱子存放的货物数AM1, M2, M3, M4, ...
输入为二行,格式如下:
N K
MI M2 M3 M4
N和K的取值范围皆为[1, 100000]
第i个包装箱存放货物数的取值范围也是[1, 100000]
*/

package main

import "fmt"

func main() {
	var (
		n, k    int
		tmp     int
		current int
		count   []int
		result  int
	)
	fmt.Scanln(&n, &k)
	count = make([]int, k)
	// 输入货物列表，实时累加
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		current = (current + tmp) % k
		count[current] += 1
	}
	for _, v := range count {
		if v > 1 {
			result += v * (v - 1) / 2
		}
	}
	fmt.Println(result)
}
