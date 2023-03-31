package storage

import (
	"github.com/ARdekS/homework-qs/internal/tree"
)

type Storage struct {
	Storage map[int]*tree.Node
	Head    int
}

func NewStorage() *Storage {
	return &Storage{Storage: make(map[int]*tree.Node, 0), Head: 0}
}

func (s *Storage) AddItem(n tree.Node) []int {
	var deletedItems []int
	//Node already in Storage
	if i, ok := s.Storage[n.ID]; ok {
		if n.IsDeleted {
			deletedItems = s.DeleteItem(n.ID)
		}
		i.Edit(n)

	} else { //Node is new in Storage
		m := n.Copy()
		s.Storage[m.ID] = &m
		if n.ParentID != 0 {
			s.Storage[m.ParentID].Nodes = append(s.Storage[m.ParentID].Nodes, m.ID)
		}

	}
	return deletedItems

}

func (s *Storage) InitTree() {
	tree.ResetID()
	data := tree.Node{
		Text: "My TODO list", ID: tree.GenerateID(), ParentID: 0, Nodes: make([]int, 0)}
	s.Head = data.ID
	s.AddItem(data)

	for i := 0; i < 3; i++ {
		s.AddItem(s.Storage[1].AddChild())
	}
	for i := 0; i < 3; i++ {
		s.AddItem(s.Storage[2].AddChild())
		s.AddItem(s.Storage[3].AddChild())
	}
	s.AddItem(s.Storage[6].AddChild())
	s.AddItem(s.Storage[7].AddChild())
	s.AddItem(s.Storage[8].AddChild())

}

func (s *Storage) GetItem(id int) *tree.Node {
	return s.Storage[id]
}

func (s *Storage) DeleteItem(id int) []int {
	deletedItems := make([]int, 0)
	if n, ok := s.Storage[id]; ok {
		if !n.IsDeleted {
			n.IsDeleted = true
			deletedItems = append(deletedItems, id)
			for _, i := range n.Nodes {
				deletedItems = append(deletedItems, s.DeleteItem(i)...)
			}
		}
	}
	return deletedItems
}
