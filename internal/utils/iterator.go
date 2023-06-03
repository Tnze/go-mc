package utils

import "math"

// Iterator represents an iterator over a collection of values.
type Iterator[T any] struct {
	data *[]T
}

// NewIterator creates a new Iterator with the given data.
func NewIterator[T any](data *[]T) *Iterator[T] {
	return &Iterator[T]{data}
}

// Filter returns a new Iterator with values that satisfy the given predicate function.
func (it *Iterator[T]) Filter(predicate func(T) bool) *Iterator[T] {
	filteredData := make([]T, 0)
	for _, value := range *it.data {
		if predicate(value) {
			filteredData = append(filteredData, value)
		}
	}
	return &Iterator[T]{&filteredData}
}

// OrderBy returns a new Iterator with values that are sorted by the given comparator function.
func (it *Iterator[T]) OrderBy(comparator func(T, T) int) *Iterator[T] {
	orderedData := make([]T, len(*it.data))
	copy(orderedData, *it.data)
	quickSort(orderedData, comparator)
	return &Iterator[T]{&orderedData}
}

// MinOrNull returns the minimum value in the Iterator or nil if the Iterator is empty.
func (it *Iterator[T]) MinOrNull(comparator func(T, T) int) *T {
	min := math.MinInt64
	if len(*it.data) == 0 {
		return nil
	}
	minValue := (*it.data)[0]
	for _, value := range *it.data {
		if comparator(value, minValue) < min {
			min = comparator(value, minValue)
			minValue = value
		}
	}
	return &minValue
}

// MaxOrNull returns the maximum value in the Iterator or nil if the Iterator is empty.
func (it *Iterator[T]) MaxOrNull(comparator func(T, T) int) *T {
	max := math.MaxInt64
	if len(*it.data) == 0 {
		return nil
	}
	maxValue := (*it.data)[0]
	for _, value := range *it.data {
		if comparator(value, maxValue) > max {
			max = comparator(value, maxValue)
			maxValue = value
		}
	}
	return &maxValue
}

// Count returns the number of values in the Iterator.
func (it *Iterator[T]) Count() int {
	return len(*it.data)
}

// ForEach executes the given function for each value in the Iterator.
func (it *Iterator[T]) ForEach(action func(T)) {
	for _, value := range *it.data {
		action(value)
	}
}

func quickSort[T any](data []T, comparator func(T, T) int) {
	if len(data) <= 1 {
		return
	}
	pivot := data[len(data)-1]
	i := 0
	for j := 0; j < len(data)-1; j++ {
		if comparator(data[j], pivot) <= 0 {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[len(data)-1] = data[len(data)-1], data[i]
	quickSort(data[:i], comparator)
	quickSort(data[i+1:], comparator)
}
