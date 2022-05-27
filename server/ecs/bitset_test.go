package ecs

import (
	"reflect"
	"testing"
)

func TestBitSet_And(t *testing.T) {
	var set1, set2 BitSet

	set1.Set(1)
	set1.Set(3)
	set1.Set(40)

	set2.Set(2)
	set2.Set(3)
	set2.Set(9)
	set2.Set(40)

	var results []Index
	set1.And(&set2).Range(func(eid Index) {
		results = append(results, eid)
	})
	want := []Index{3, 40}
	if !reflect.DeepEqual(results, want) {
		t.Errorf("want %v, got: %v", want, results)
	}
}
