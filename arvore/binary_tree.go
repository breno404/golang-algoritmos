package arvore

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) Insert(value int) {
	t.root = insertNode(t.root, value)
}

func insertNode(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}

	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else if value > node.Value {
		node.Right = insertNode(node.Right, value)
	}

	return node
}

func (t *BinaryTree) Search(value int) *Node {
	return searchNode(t.root, value)
}

func searchNode(node *Node, value int) *Node {

	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}

	if value < node.Value {
		return searchNode(node.Left, value)
	}

	return searchNode(node.Right, value)

}

func (t *BinaryTree) Delete(value int) {
	t.root = deleteNode(t.root, value)
}

func deleteNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		min := findMin(node.Right)

		node.Value = min.Value

		node.Right = deleteNode(node.Right, min.Value)
	}

	return node
}

func findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

func (t *BinaryTree) Inverse() {
	inverseTree(t.root)
}

func inverseTree(node *Node) {
	if node == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left
	inverseTree(node.Left)
	inverseTree(node.Right)
}
