package arvore

type AVLTree struct {
	root *NodeWithHeight
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func height(n *NodeWithHeight) int {
	if n == nil {
		return 0
	}

	return n.Height
}

func getBalance(n *NodeWithHeight) int {
	if n == nil {
		return 0
	}

	return height(n.Left) - height(n.Right)
}

func rotateLeft(x *NodeWithHeight) *NodeWithHeight {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	return y
}

func rotateRight(y *NodeWithHeight) *NodeWithHeight {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	return x
}

func findMinAVL(node *NodeWithHeight) *NodeWithHeight {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

func (t *AVLTree) Insert(value int) *NodeWithHeight {
	t.root = insertAVL(t.root, value)

	return t.root
}

func insertAVL(node *NodeWithHeight, value int) *NodeWithHeight {
	if node == nil {
		return &NodeWithHeight{Value: value, Height: 1}
	}

	if value < node.Value {
		node.Left = insertAVL(node.Left, value)
	} else if value > node.Value {
		node.Right = insertAVL(node.Right, value)
	} else {
		return node
	}

	node.Height = max(height(node.Left), height(node.Right)) + 1

	balance := getBalance(node)

	if balance > 1 && value < node.Left.Value {
		return rotateRight(node)
	}

	if balance < -1 && value > node.Right.Value {
		return rotateLeft(node)
	}

	if balance > 1 && value > node.Left.Value {
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	}

	if balance < -1 && value < node.Right.Value {
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}

	return node
}

func (t *AVLTree) Search(value int) *NodeWithHeight {
	return searchAVL(t.root, value)
}

func searchAVL(node *NodeWithHeight, value int) *NodeWithHeight {

	if value < node.Value {
		return searchAVL(node.Left, value)
	} else if value > node.Value {
		return searchAVL(node.Right, value)
	}

	return node
}

func (t *AVLTree) Delete(value int) {
	t.root = deleteAVL(t.root, value)
}

func deleteAVL(node *NodeWithHeight, value int) *NodeWithHeight {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = deleteAVL(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteAVL(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		temp := findMinAVL(node.Right)
		node.Value = temp.Value
		node.Right = deleteAVL(node.Right, temp.Value)
	}

	node.Height = max(height(node.Left), height(node.Right)) + 1

	balance := getBalance(node)

	if balance > 1 && getBalance(node.Left) >= 0 {
		return rotateRight(node)
	}

	if balance > 1 && getBalance(node.Left) < 0 {
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	}

	if balance < -1 && getBalance(node.Right) <= 0 {
		return rotateLeft(node)
	}

	if balance < -1 && getBalance(node.Right) > 0 {
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}

	return node
}
