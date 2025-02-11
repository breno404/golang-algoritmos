package arvore

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

type NodeWithHeight struct {
	Left   *NodeWithHeight
	Right  *NodeWithHeight
	Value  int
	Height int
}
