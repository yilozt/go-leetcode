package main

func verifyPostorder(postorder []int) bool {
	var check func(int, int) bool
	check = func(i, j int) bool {
		// j表示树的根节点位置
		if i >= j {
			return true
		}
		tmp := i
		for postorder[tmp] < postorder[j] {
			tmp++
		}
		mid := tmp
		for postorder[tmp] > postorder[j] {
			tmp++
		}
		return tmp == j && check(i, mid-1) && check(mid, j-1)
	}
	return check(0, len(postorder)-1)
}
