package data

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type HashTable[T, K constraints.Unsigned] struct {
	data map[T]map[K]int
}

func NewHashTable[T, K constraints.Unsigned]() *HashTable[T, K] {
	return &HashTable[T, K]{make(map[T]map[K]int)}
}

func (h *HashTable[T, K]) Put(rowKey T, columnKey K, value int) {
	if h.data[rowKey] == nil {
		h.data[rowKey] = make(map[K]int)
	}
	h.data[rowKey][columnKey] = value
}

func (h *HashTable[T, K]) Get(rowKey T, columnKey K) int {
	row, ok := h.data[rowKey]
	if !ok {
		fmt.Println(fmt.Errorf("rowKey %v not found", rowKey))
		return 0
	}
	value, _ := row[columnKey]
	return value
}
