package msg

import (
	"crypto/rand"
	"testing"

	"github.com/Tnze/go-mc/chat/sign"
)

func TestSignatureCache(t *testing.T) {
	cache := newSignatureCache()
	s1 := &sign.Signature{1}
	s2 := &sign.Signature{2}
	s3 := &sign.Signature{3}
	s4 := &sign.Signature{4}
	// t.Logf("%p, %p, %p, %p", s1, s2, s3, s4)
	err := cache.popOrInsert(nil, []sign.PackedSignature{
		{Signature: s1},
		{Signature: s2},
		{Signature: s3},
	})
	if err != nil {
		t.Fatal(err)
	}
	// cache: [s1, s2, s3, nil...]
	if cache.signatures[0] != s1 || cache.signatures[1] != s2 || cache.signatures[2] != s3 {
		t.Log(cache.signatures)
		t.Fatal("insert s1~3 error")
	}
	err = cache.popOrInsert(s4, []sign.PackedSignature{{Signature: s3}})
	if err != nil {
		t.Fatal(err)
	}
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
	cache := newSignatureCache()
	signs := make([]sign.PackedSignature, len(cache.signatures)+5)
	for i := range signs {
		newSign := new(sign.Signature)
		_, _ = rand.Read(newSign[:])
		signs[i] = sign.PackedSignature{Signature: newSign}
	}
	err := cache.popOrInsert(nil, signs[:len(cache.signatures)])
	if !signatureEquals(cache.signatures[:], signs[:len(cache.signatures)]) {
		t.Fatal("insert error")
	}
	if err != nil {
		t.Fatal(err)
	}
	insert2 := signs[len(cache.signatures)-5:]
	err = cache.popOrInsert(nil, insert2)
	if err != nil {
		t.Fatal(err)
	}
	if !signatureEquals(cache.signatures[:10], insert2) {
		t.Fatal("insert and pop error")
	}
}

func signatureEquals(a []*sign.Signature, b []sign.PackedSignature) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i].Signature {
			return false
		}
	}
	return true
}
