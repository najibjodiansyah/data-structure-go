package ds

import "fmt"

type HashTable struct {
	data []any
}

type HashTableInterface interface {
	Set(key any, value any)
	Get(key any) any
	Keys() []any
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		data: make([]any, size),
	}
}

func (h *HashTable) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		fmt.Println(key[i], i, string(key[i]))
		hash = (hash + int(key[i])*i) % len(h.data)
	}
	fmt.Println("hash of ", key, "is", hash)
	return hash
}

func (h *HashTable) Set(key any, value any) {
	address := h.hash(key.(string))
	if h.data[address] == nil {
		h.data[address] = make([][]any, 0)
	}
	h.data[address] = append(h.data[address].([][]any), []any{key, value})
}

func (h *HashTable) Get(key any) any {
	address := h.hash(key.(string))
	bucket := h.data[address].([][]any)
	for i := 0; i < len(bucket); i++ {
		if bucket[i][0] == key {
			return bucket[i][1]
		}
	}
	return nil
}
