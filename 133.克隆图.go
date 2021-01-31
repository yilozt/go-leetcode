/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

package main

func cloneGraph(node *Node) *Node {
	visited := make(map[int]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[int]*Node) *Node {
	// 节点为空
	if node == nil {
		return nil
	}
	// 已经克隆过该节点
	if cloned, ok := visited[node.Val]; ok {
		return cloned
	}
	// 没有克隆过该节点，首先分配空间
	result := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	// 记录为已克隆
	visited[node.Val] = result
	// 克隆所有的边
	for index, neighbor := range node.Neighbors {
		result.Neighbors[index] = clone(neighbor, visited)
	}
	return result
}

// @lc code=end
