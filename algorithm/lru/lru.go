package lru

type LRUCache struct {
	size       int
	capacity   int
	cache      map[string]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value string
	prev, next *DLinkedNode
}

func NewLinkedNode(key, value string) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func NewLRUCache(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    map[string]*DLinkedNode{},
		head:     NewLinkedNode("", ""),
		tail:     NewLinkedNode("", ""),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key string) string {
	if _, ok := l.cache[key]; !ok {
		return ""
	}

	node := l.cache[key]
	l.moveToHead(node)
	return node.value
}

func (l *LRUCache) Put(key, value string) {
	if _, ok := l.cache[key]; !ok {
		node := NewLinkedNode(key, value)
		l.cache[key] = node
		l.addToHead(node)
		l.size++

		if l.size > l.capacity {
			removed := l.removeToTail()
			delete(l.cache, removed.key)
		}
	} else {
		node := l.cache[key]
		node.value = value
		l.moveToHead(node)
	}
}

func (l *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) moveToHead(node *DLinkedNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeToTail() *DLinkedNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}
