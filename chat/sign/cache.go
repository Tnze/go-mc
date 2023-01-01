package sign

import (
	"errors"
)

type SignatureCache struct {
	signatures   [128]*Signature
	signIndexes  map[Signature]int
	cachedBuffer []*Signature
}

func NewSignatureCache() SignatureCache {
	return SignatureCache{
		signIndexes: make(map[Signature]int),
	}
}

func (s *SignatureCache) PopOrInsert(self *Signature, lastSeen []*Signature) {
	var tmp *Signature
	s.cachedBuffer = s.cachedBuffer[:0] // clear buffer
	if self != nil {
		s.cachedBuffer = append(s.cachedBuffer, self)
	}
	for _, v := range lastSeen {
		s.cachedBuffer = append(s.cachedBuffer, v)
	}
	for i := 0; i < len(s.cachedBuffer) && i < len(s.signatures); i++ {
		v := s.cachedBuffer[i]
		if i, ok := s.signIndexes[*v]; ok {
			s.signatures[i] = nil
		}
		tmp, s.signatures[i] = s.signatures[i], v
		s.signIndexes[*v] = i
		if tmp != nil {
			s.cachedBuffer = append(s.cachedBuffer, tmp)
			delete(s.signIndexes, *tmp)
		}
	}
}

var UncachedSignature = errors.New("uncached signature")
