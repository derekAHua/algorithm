package internal

import "container/list"

// @Author: Derek
// @Description:
// @Date: 2023/8/31 22:12
// @Version 1.0

type keyLru struct {
	limit    int                      //缓存数量
	evicts   *list.List               //双向链表用于淘汰数据
	elements map[string]*list.Element //记录缓存数据

	onEvict func(key string)
}

func NewKeyLru(limit int, onEvict func(key string)) *keyLru {
	return &keyLru{
		limit:    limit,
		evicts:   list.New(),
		elements: make(map[string]*list.Element),
		onEvict:  onEvict,
	}
}

func (l *keyLru) add(key string) {
	if v, ok := l.elements[key]; ok {
		l.evicts.MoveToFront(v)
		return
	}

	elem := l.evicts.PushFront(key)
	l.elements[key] = elem

	if l.evicts.Len() > l.limit {
		l.removeOldest()
	}

	return
}

func (l *keyLru) removeOldest() {
	elem := l.evicts.Back() //获取链表末尾节点
	if elem != nil {
		l.evicts.Remove(elem)
		key := elem.Value.(string)
		delete(l.elements, key)
		l.onEvict(key)
	}
}
