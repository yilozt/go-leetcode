/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func pathSum(root *TreeNode, sum int) [][]int {
	// 基本dfs
	var result [][]int
	if root == nil {
		return result
	}
	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, current int, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		current += node.Val
		// 符合条件，添加结果
		if current == sum && node.Left == nil && node.Right == nil {
			// 添加的结果需要进行深拷贝
			result = append(result, append([]int(nil), path...))
		} else {
			dfs(node.Left, current, path)
			dfs(node.Right, current, path)
		}
		// 最后需要去掉路径中的最后一个值
		path = path[:len(path)-1]
	}
	dfs(root, 0, []int{})
	return result
}
