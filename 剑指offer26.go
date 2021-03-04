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
	if A.Val == B.Val && fixedSubStructure(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func fixedSubStructure(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return fixedSubStructure(A.Left, B.Left) && fixedSubStructure(A.Right, B.Right)
}
