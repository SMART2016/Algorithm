package main

import "container/list"

type LRU struct {
	l        list.List
	m        map[int]int
	capacity int
}

func (c *LRU) New(capacity int) {
	//	c.l = list.New()
	//	c.m = make(map[int]int)
	//	c.capacity = capacity
}

func (c *LRU) Get(key *list.Element) int {

	if val, ok := c.m[key.Value.(int)]; ok {
		c.l.Remove(key)
		c.l.PushFront(key)
		return val
	}
	return -1
}

func (c *LRU) Set(key int, value int) {
	//	if val, ok := c.m[key]; ok {

	c.m[key] = value
	//		c.l.Remove(key)
	//}
}
