package sign

import (
	"crypto/rand"
	"testing"
)

func TestSignatureCache(t *testing.T) {
	cache := NewSignatureCache()
	s1 := &Signature{1}
	s2 := &Signature{2}
	s3 := &Signature{3}
	s4 := &Signature{4}
	// t.Logf("%p, %p, %p, %p", s1, s2, s3, s4)
	cache.PopOrInsert(nil, []*Signature{s1, s2, s3})
	// cache: [s1, s2, s3, nil...]
	if cache.signatures[0] != s1 || cache.signatures[1] != s2 || cache.signatures[2] != s3 {
		t.Log(cache.signatures)
		t.Fatal("insert s1~3 error")
	}
	cache.PopOrInsert(s4, []*Signature{s3})
	// cache: [s4, s3, s1, s2, nil...]
	if cache.signatures[0] != s4 {
		t.Log(cache.signatures)
		t.Fatal("insert s4 error")
	}
	if cache.signatures[1] != s3 {
		t.Log(cache.signatures)
		t.Fatal("pop s3 error")
	}
	if cache.signatures[2] != s1 || cache.signatures[3] != s2 {
		t.Log(cache.signatures)
		t.Fatal("s1~2 position error")
	}
}

func TestSignatureCache_2(t *testing.T) {
	cache := NewSignatureCache()
	signs := make([]*Signature, len(cache.signatures)+5)
	for i := range signs {
		signs[i] = new(Signature)
		_, _ = rand.Read(signs[i][:])
	}
	cache.PopOrInsert(nil, signs[:len(cache.signatures)])
	if !signatureEquals(cache.signatures[:], signs[:len(cache.signatures)]) {
		t.Fatal("insert error")
	}
	insert2 := signs[len(cache.signatures)-5:]
	cache.PopOrInsert(nil, insert2)
	if !signatureEquals(cache.signatures[:10], insert2) {
		t.Fatal("insert and pop error")
	}
}

func signatureEquals(a, b []*Signature) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
