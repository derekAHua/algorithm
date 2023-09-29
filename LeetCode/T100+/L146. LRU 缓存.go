package T100_

type LRUCacheNode struct {
	Val  int
	Pre  *LRUCacheNode
	Next *LRUCacheNode
}

type LRUCache struct {
	capacity int
	cnt      int
	m        map[int]*LRUCacheNode
	head     *LRUCacheNode
	tail     *LRUCacheNode
}

func Constructor(capacity int) (r LRUCache) {
	r.capacity = capacity
	r.m = make(map[int]*LRUCacheNode)
	r.head = new(LRUCacheNode)
	r.tail = new(LRUCacheNode)
	r.head.Next = r.tail
	r.tail.Pre = r.head
	return
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.m[key]
	if !ok {
		return -1
	}
	return val.Val
}

func (this *LRUCache) Put(key int, value int) {
	if this.cnt == this.capacity {
		rm := this.tail.Pre
		rm.Pre.Next = this.tail
		this.tail.Pre = rm.Pre
		rm = nil
		this.cnt--
	}

	cur := new(LRUCacheNode)
	cur.Val = value

	this.m[key] = cur

	this.head.Pre = cur
	cur.Next = this.head
	this.head = cur

	this.cnt++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
