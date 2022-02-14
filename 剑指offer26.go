/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	if A.Val == B.Val && fixedSubStructure(A, B) { // 根相同则继续比较子树
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B) //否则递归比较B和A的左右子树
}

func fixedSubStructure(A, B *TreeNode) bool {
	if B == nil { // B为空，说明A包含了B
		return true
	}
	if A == nil || A.Val != B.Val { // A为空或者A和B的值不等，说明未包含B
		return false
	}
	return fixedSubStructure(A.Left, B.Left) && fixedSubStructure(A.Right, B.Right) // 递归比较A和B的左右子树
}
