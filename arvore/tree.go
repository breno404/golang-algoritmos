package arvore

type Tree interface {
	Insert(int)
	Search(int) *Node
	Delete(int)
	Inverse()
	PreOrder()
	InOrder()
	PostOrder()
	LevelOrder()
}
