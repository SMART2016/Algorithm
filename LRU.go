package main

import "container/list"

type LRU struct {
	l        *list.List
	m        map[int]int
	capacity int
}

func (c *LRU) New(capacity int) {
	c.l = list.New()
	c.m = make(map[int]int)
	c.capacity = capacity
}

func (c *LRU) Get(key int) int {

	if val, ok := c.m[key]; ok {
		c.adjustElement(key)
		return val
	}
	return -1
}

func (c *LRU) Set(key int, value int) {
	c.m[key] = value
	if c.l.Len() > c.capacity {
		removeKey := c.l.Remove(c.l.Front().Prev())
		delete(c.m, removeKey.(int))
	}
	c.adjustElement(key)
}

func (c *LRU) find(l *list.List, elem list.Element) *list.Element {
	for head := l.Front(); head != nil; {
		if head.Value == elem.Value {
			return head
		} else {
			head = head.Next()
		}
	}
	return nil
}

func (c *LRU) adjustElement(val int) {
	listElement := list.Element{Value: val}
	listNode := c.find(c.l, listElement)
	if listNode != nil {
		c.l.Remove(listNode)
	}
	c.l.PushFront(val)
}
