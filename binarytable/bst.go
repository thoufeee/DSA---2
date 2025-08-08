package binarytable

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// type BST struct {
// 	Root *Node
// }

// // insert
// func (bst *BST) Insert(val int) {
// 	bst.Root = InsertNode(bst.Root, val)
// }

// // insert node
// func InsertNode(node *Node, val int) *Node {
// 	if node == nil {
// 		return &Node{data: val}
// 	}

// 	if val < node.data {
// 		node.left = InsertNode(node.left, val)
// 	} else if val > node.data {
// 		node.right = InsertNode(node.right, val)
// 	}

// 	return node
// }

// // search
// func (bst *BST) Search(val int) bool {
// 	return SearchNode(bst.Root, val)
// }

// // search node
// func SearchNode(node *Node, val int) bool {
// 	if node == nil {
// 		return false
// 	}

// 	if val == node.data {
// 		return true
// 	} else if val < node.data {
// 		return SearchNode(node.left, val)
// 	} else {
// 		return SearchNode(node.right, val)
// 	}
// }

// // delete
// func (bst *BST) Delete(val int) {
// 	bst.Root = DeleteNode(bst.Root, val)
// }

// // delete node
// func DeleteNode(node *Node, val int) *Node {
// 	if node == nil {
// 		return nil
// 	}

// 	if val < node.data {
// 		node.left = DeleteNode(node.left, val)
// 	} else if val > node.data {
// 		node.right = DeleteNode(node.right, val)
// 	} else {

// 		if node.left == nil {
// 			return node.right
// 		}

// 		if node.right == nil {
// 			return node.left
// 		}

// 		minright := Min(node.right)
// 		node.data = minright.data
// 		node.right = DeleteNode(node.right, minright.data)

// 	}

// 	return node
// }

// // min
// func Min(node *Node) *Node {
// 	current := node

// 	for current.left != nil {
// 		current = current.left
// 	}

// 	return current

// }

// // traversals

// // inorder
// func Inorder(node *Node) {
// 	if node != nil {
// 		Inorder(node.left)
// 		fmt.Print(node.data, " ")
// 		Inorder(node.right)

// 	}
// }

// // preorder
// func PreOrder(node *Node) {
// 	if node != nil {
// 		fmt.Print(node.data, " ")
// 		PreOrder(node.left)
// 		PreOrder(node.right)
// 	}
// }

// // postorder
// func PostOrder(node *Node) {
// 	if node != nil {
// 		PostOrder(node.left)
// 		PostOrder(node.right)
// 		fmt.Print(node.data, " ")
// 	}
// }
