package main

import "container/list"

type LRUCache struct {
	capacity  int
	evictList *list.List
	items     map[interface{}]*list.Element
}

type entry struct {
	key   interface{}
	value interface{}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		evictList: list.New(),
		items:     make(map[interface{}]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if ent, ok := this.items[key]; ok {
		this.evictList.MoveToFront(ent)
		if ent.Value.(*entry) == nil {
			return -1
		}
		return ent.Value.(*entry).value.(int)
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ent, ok := this.items[key]; ok {
		this.evictList.MoveToFront(ent)
		ent.Value.(*entry).value = value
		return
	}
	ent := &entry{
		key:   key,
		value: value,
	}
	entry := this.evictList.PushFront(ent)
	this.items[key] = entry

	if this.evictList.Len() > this.capacity {
		this.removeOldest()
	}
}

func (this *LRUCache) removeOldest() {
	backNode := this.evictList.Back()
	delete(this.items, backNode.Value.(*entry).key)
	this.evictList.Remove(backNode)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
