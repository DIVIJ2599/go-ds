package hashtable

type HashTable struct {
	array []*bucket
	size  int
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func New(size int) *HashTable {
	if size <= 0 {
		size = 7
	}

	result := &HashTable{
		array: make([]*bucket, size),
		size:  size,
	}

	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func (h *HashTable) Insert(key string) {
	index := hash(key, h.size)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key, h.size)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key, h.size)
	h.array[index].delete(key)
}

func hash(key string, size int) int {
	hash := 0

	for _, v := range key {
		hash += int(v)
	}

	return hash % size
}

func (b *bucket) insert(key string) {
	if b.search(key) {
		return
	}

	newNode := &bucketNode{key: key}
	newNode.next = b.head
	b.head = newNode
}

func (b *bucket) search(key string) bool {
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == key {
			return true
		}

		currentNode = currentNode.next
	}

	return false
}

func (b *bucket) delete(key string) {
	if b.head == nil {
		return
	}

	if b.head.key == key {
		b.head = b.head.next
		return
	}

	previousNode := b.head

	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
			return
		}

		previousNode = previousNode.next
	}
}
