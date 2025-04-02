package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

var idx int = -1

func buildTree(preorder []int) *Node {
	idx++
	if preorder[idx] == -1 {
		return nil
	}
	root := &Node{data: preorder[idx]}
	root.left = buildTree(preorder)
	root.right = buildTree(preorder)
	return root
}

func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.data)
	fmt.Print(" ")
	PreOrderTraversal(root.left)
	PreOrderTraversal(root.right)
}

func main() {
	preorder := []int{1, 2, -1, -1, 3, 4, -1, -1, 5, -1, -1}
	root := buildTree(preorder)
	PreOrderTraversal(root)
}
