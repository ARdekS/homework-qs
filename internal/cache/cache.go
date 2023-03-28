package cache

import (
	"github.com/ARdekS/homework-qs/internal/tree"
)

type Cache struct {
	Cache map[int]*tree.Node
	Heads []int // содержить в себе все корневые ноды для кэша
}

func NewCache() *Cache {
	return &Cache{
		Cache: make(map[int]*tree.Node, 0),
	}
}

func (c *Cache) GetNode(id int) *tree.Node {
	return c.Cache[id]
}

func (c *Cache) AddItem(n *tree.Node) {
	if _, ok := c.Cache[n.ID]; !ok {
		m := n.Copy()
		c.Cache[m.ID] = &m
		//если n это чей либо родитель то мы должны добавить детей и скрыть их
		for _, i := range n.Nodes {
			if child, ok := c.Cache[i]; ok {
				m.Nodes = append(m.Nodes, child.ID)
				for i, id := range c.Heads {
					if id == child.ID {
						c.Heads = append(c.Heads[:i], c.Heads[i+1:]...)
					}
				}
			}
		}
		//если родитель для n уже есть то мы должны добавить ему ребенка n
		if p, ok := c.Cache[m.ParentID]; ok {
			p.Nodes = append(p.Nodes, m.ID)
			if p.IsDeleted {
				c.DeleteItem(m.ID)
			}
		} else {
			c.Heads = append(c.Heads, m.ID)
		}
	}
}
func (c *Cache) NewItme(m tree.Node) {
	n := c.GetNode(m.ID)
	child := n.AddChild()
	c.AddItem(&child)

}
func (c *Cache) EditItem(m tree.Node) {
	if n, ok := c.Cache[m.ID]; ok {
		n.Text = m.Text
	}
}
func (c *Cache) DeleteItem(id int) {
	if n, ok := c.Cache[id]; ok {
		n.IsDeleted = true
		for _, i := range n.Nodes {
			c.DeleteItem(i)
		}
	}
}

func (c *Cache) ReturnItems() []int {
	// var result []int = make([]int, 0, 1)
	// for k, v := range c.Cache {
	// 	if !v.IsLinked {
	// 		result = append(result, k)
	// 	}
	// }
	// sort.SliceStable(result, func(i, j int) bool {
	// 	return result[i] < result[j]
	// })
	// return result
	return c.Heads
}
