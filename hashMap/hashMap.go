package hashmap

import (
	"hash/fnv"
)

type Node struct {
	key   string
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(key string, value int) {
	node := &Node{key: key, value: value, next: l.head}
	l.head = node
}

func (l *LinkedList) Get(key string) (int, bool) {
	current := l.head
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

type HashMap struct {
	buckets []*LinkedList
	size    int
}

func NewHashMap(size int) *HashMap {
	buckets := make([]*LinkedList, size)
	for i := range buckets {
		buckets[i] = &LinkedList{}
	}
	return &HashMap{buckets: buckets, size: size}
}

func (h *HashMap) hash(key string) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	return int(hasher.Sum32()) % h.size
}

func (h *HashMap) Insert(key string, value int) {
	index := h.hash(key)
	h.buckets[index].Insert(key, value)
}

func (h *HashMap) Get(key string) (int, bool) {
	index := h.hash(key)
	return h.buckets[index].Get(key)
}
