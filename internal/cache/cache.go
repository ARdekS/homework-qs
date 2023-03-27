package cache

import (
	"sort"

	"github.com/ARdekS/homework-qs/internal/tree"
)

// type ID struct {
// 	uint64 `json:"id"`
// }

type Cache struct {
	Cache map[uint64]*tree.Node
	// Parents map[uint64]*tree.Node
}

func NewCache() *Cache {
	return &Cache{Cache: make(map[uint64]*tree.Node, 0)}
}

func (c *Cache) AddItem(n *tree.Node) {
	if _, ok := c.Cache[n.ID]; !ok {
		//если n это чей либо родитель то мы должны добавить детей и скрыть их
		for _, i := range c.Cache {
			if i.ParentID == n.ID {
				n.Nodes = append(n.Nodes, i)
				i.IsChild = true
			}
		}
		//если родитель для n уже есть то мы должны добавить ему ребенка n
		if p, ok := c.Cache[n.ParentID]; ok {
			n.IsChild = true
			p.Nodes = append(p.Nodes, n)
		}

		c.Cache[n.ID] = n

	}
}
func (c *Cache) EditItem(m *tree.Node) {
	if n, ok := c.Cache[m.ID]; ok {
		n.Merge(m)
	}
}

func (c *Cache) ReturnItems() []*tree.Node {
	var result []*tree.Node = make([]*tree.Node, 0, 1)
	for _, v := range c.Cache {
		if !v.IsChild {
			result = append(result, v)
		}
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result
}
