package msg

import (
	"github.com/Tnze/go-mc/chat/sign"
)

type signatureCache struct {
	signatures   [128]*sign.Signature
	signIndexes  map[sign.Signature]int
	cachedBuffer []*sign.Signature
}

func newSignatureCache() signatureCache {
	return signatureCache{
		signIndexes: make(map[sign.Signature]int),
	}
}

func (s *signatureCache) popOrInsert(self *sign.Signature, lastSeen []sign.PackedSignature) error {
	var tmp *sign.Signature
	s.cachedBuffer = s.cachedBuffer[:0] // clear buffer
	if self != nil {
		s.cachedBuffer = append(s.cachedBuffer, self)
	}
	for _, v := range lastSeen {
		if v.Signature != nil {
			s.cachedBuffer = append(s.cachedBuffer, v.Signature)
		} else if v.ID >= 0 && int(v.ID) < len(s.signatures) {
			s.cachedBuffer = append(s.cachedBuffer, s.signatures[v.ID])
		} else {
			return InvalidChatPacket
		}
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
	return nil
}
