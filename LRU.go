package main

import "container/list"

type LRU struct {
	l        List
	m        map[int]int
	capacity int
}

func (c *LRU) New(capacity int) {
	c.l = list.New()
	c.m = make(map[int]int)
	c.capacity = capacity
}

func (c *LRU) Get(key int) int {

}

func (c *LRU) Set(key int, value int) {

}
