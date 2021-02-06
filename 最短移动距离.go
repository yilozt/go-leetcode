package main

import (
	"fmt"
)

const (
	// 路径长度和目标节点的最大值
	MAX = 100005
)

// 存储最近有房间子节点的距离和终点
type pair struct {
	distance int
	goal     int
}

var (
	// 节点房间数
	nodes = make([]int, MAX)
	// 存储每个节点到最近有房间子节点的终点和距离
	cache = make([]pair, MAX)
	// 存储该松鼠到有房间节点的最好情况
	best pair
	// 节点和松鼠的个数
	n, m int
	// 公共根
	public int
	// 方向
	direction = make([]int, MAX)
)

func main() {
	fmt.Scan(&n)
	fmt.Scan(&m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&nodes[i])
	}
	fmt.Println(solution())
}

func adjust(a, b int) int {
	if a*b < 0 {
		return -1
	}
	return 1
}

// 更新i位置的节点cache信息
func update(i int) {
	if nodes[i] > 0 {
		// 距离为0,终点为自己
		cache[i] = pair{0, i}
	} else {
		// 初始化为最大
		cache[i] = pair{MAX, MAX}
	}
	// 查找子节点房间
	j := i * 2
	// 左子节点
	if j <= n && cache[j].distance+adjust(direction[j], 1) < cache[i].distance {
		cache[i] = pair{cache[j].distance + adjust(direction[j], 1), cache[j].goal}
	}
	// 右子节点
	j += 1
	if j <= n && cache[j].distance+adjust(direction[j], 1) < cache[i].distance {
		cache[i] = pair{cache[j].distance + adjust(direction[j], 1), cache[j].goal}
	}
}

func queryForBest(i, cost int) {
	// 尝试逐级向上走再向下走，确认是否有更近情况
	if cache[i].distance+cost < best.distance {
		public = i
		best = pair{cache[i].distance + cost, cache[i].goal}
	}
	if i != 1 {
		queryForBest(i/2, cost+adjust(direction[i], -1))
	}
}

func goUp(i int) {
	update(i)
	if i == public {
		return
	}
	direction[i] -= 1
	goUp(i / 2)
}

func goDown(i int) {
	update(i)
	if i == public {
		return
	}
	direction[i] += 1
	goDown(i / 2)
}

func solution() int {
	var result int
	// 从后往前更新节点缓存
	for i := n; i >= 1; i-- {
		update(i)
	}
	// 遍历松鼠
	var squ int
	for i := 1; i <= m; i++ {
		best = pair{MAX, MAX}
		fmt.Scan(&squ)
		// 找到最佳路径并存储到best中
		queryForBest(squ, 0)
		// 结果增加
		result += best.distance
		goUp(squ)
		// 目标节点房间数减少
		nodes[best.goal] -= 1
		goDown(best.goal)
		for ; squ != 0; squ >>= 1 {
			update(squ)
		}
	}
	return result
}
