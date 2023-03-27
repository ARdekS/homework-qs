package tree

import "fmt"

type Storage struct {
	Storage map[uint64]*Node
	Head    *Node
}

func NewStorage() *Storage {
	return &Storage{Storage: make(map[uint64]*Node, 0), Head: nil}
}

func (s *Storage) AddToStorage(n *Node) {
	//add head to storage
	var m = n.Copy()
	if s.Head == nil && n.ParentID == 0 {
		s.Storage[n.ID] = m
		s.Head = m
	} else if i, ok := s.Storage[n.ID]; ok {
		i.Merge(n)
	} else { //new nodes
		s.Storage[n.ID] = m
		if p, ok := s.Storage[n.ParentID]; ok {
			p.Nodes = append(p.Nodes, m)
		}
	}

	// for _, child := range n.Nodes {
	// 	s.AddToStorage(child)
	// }

	// s.Storage[n.ID] = n

}

func Traverse(node *Node) {
	fmt.Println(node.Text)
	for _, child := range node.Nodes {
		Traverse(child)
	}
}

func InitTree() *Storage {
	storage := NewStorage()

	data := &Node{
		Text: "My TODO list", ID: 0, ParentID: 0}
	storage.AddToStorage(data)
	for i := 0; i < 3; i++ {
		b := data.addChild(nil)
		storage.AddToStorage(&b)
	}
	for i := 0; i < 3; i++ {
		b := data.Nodes[1].addChild(nil)
		storage.AddToStorage(&b)
		c := data.Nodes[2].addChild(nil)
		storage.AddToStorage(&c)
	}
	a := data.Nodes[2].addChild(nil)
	b := data.Nodes[1].Nodes[2].addChild(nil)
	c := data.Nodes[1].Nodes[1].addChild(nil)
	d := data.Nodes[1].Nodes[1].Nodes[0].addChild(nil)
	storage.AddToStorage(&a)
	storage.AddToStorage(&b)
	storage.AddToStorage(&c)
	storage.AddToStorage(&d)

	return storage
}
