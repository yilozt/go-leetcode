/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var pos int
	for i, v := range inorder {
		if v == preorder[0] {
			pos = i
		}
	}
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:pos+1], inorder[:pos])
	root.Right = buildTree(preorder[pos+1:], inorder[pos+1:])
	return root
}
