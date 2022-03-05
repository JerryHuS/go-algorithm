package 二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	inorder(root, &result)
	return result
}

func inorder(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, arr)
	*arr = append(*arr, node.Val)
	inorder(node.Right, arr)
}

func preorder(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	*arr = append(*arr, node.Val)
	inorder(node.Left, arr)
	inorder(node.Right, arr)
}

func postorder(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, arr)
	inorder(node.Right, arr)
	*arr = append(*arr, node.Val)
}
