package interfaces

type Node interface {
	Copy() Node
	Edit(m Node)
	AddChild() Node
}
