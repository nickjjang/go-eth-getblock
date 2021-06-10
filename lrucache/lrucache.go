package lrucache

import (
	"container/list"
)

type LRUCache struct {
	cap int
	l   *list.List
	m   map[int64]*list.Element
}

// Pair is the value of a list node.
type Pair struct {
	key   int64
	value interface{}
}

// Constructor initializes a new LRUCache.
func New(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   map[int64]*list.Element{},
	}
}

// Get a list node from the hash map.
func (c *LRUCache) Get(key int64) interface{} {
	// check if list node exists
	if node, ok := c.m[key]; ok {
		val := node.Value.(*list.Element).Value.(Pair).value
		// move node to front
		c.l.MoveToFront(node)
		return val
	}
	return false
}

// Put key and value in the LRUCache
func (c *LRUCache) Put(key int64, value interface{}) {
	// check if list node exists
	if node, ok := c.m[key]; ok {
		// move the node to front
		c.l.MoveToFront(node)
		// update the value of a list node
		node.Value.(*list.Element).Value = Pair{key: key, value: value}
	} else {
		// delete the last list node if the list is full
		if c.l.Len() == c.cap {
			// get the key that we want to delete
			itemKeyn := c.l.Back().Value.(*list.Element).Value.(Pair).key
			// delete the node pointer in the hash map by key
			delete(c.m, itemKeyn)
			// remove the last list node
			c.l.Remove(c.l.Back())
		}
		// initialize a list node
		node := &list.Element{
			Value: Pair{
				key:   key,
				value: value,
			},
		}
		// push the new list node into the list
		ptr := c.l.PushFront(node)
		// save the node pointer in the hash map
		c.m[key] = ptr
	}
}
