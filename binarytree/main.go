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

	fmt.Print(root.data, " ")
	PreOrderTraversal(root.left)
	PreOrderTraversal(root.right)
}

func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	PostOrderTraversal(root.left)
	PostOrderTraversal(root.right)
	fmt.Print(root.data, " ")
}

func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	InOrderTraversal(root.left)
	fmt.Print(root.data, " ")
	InOrderTraversal(root.right)
}

func LevelOrderTraversal(root *Node) {
	queue := []*Node{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		fmt.Print(curr.data, " ")

		if curr.left != nil {
			queue = append(queue, curr.left)
		}
		if curr.right != nil {
			queue = append(queue, curr.right)
		}
	}
}

func main() {
	preorder := []int{1, 2, -1, -1, 3, 4, -1, -1, 5, -1, -1}
	root := buildTree(preorder)
	PreOrderTraversal(root)
	fmt.Println()
	InOrderTraversal(root)
	fmt.Println()
	PostOrderTraversal(root)
	fmt.Println()
	LevelOrderTraversal(root)
}
