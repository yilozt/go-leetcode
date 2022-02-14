package main

func verifyPostorder(postorder []int) bool {
	var check func(int, int) bool
	check = func(i, j int) bool {
		// j表示树的根节点位置
		if i >= j {
			return true
		}
		tmp := i
		for postorder[tmp] < postorder[j] { // 二叉搜索树前部分值小于根节点
			tmp++
		}
		mid := tmp
		for postorder[tmp] > postorder[j] { // 后部分值大于根节点
			tmp++
		}
		return tmp == j && check(i, mid-1) && check(mid, j-1) // 首先判断该段是否符合该规则，再递归检查子树
	}
	return check(0, len(postorder)-1)
}
