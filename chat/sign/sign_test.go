package sign

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"math/big"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMessageBody_Hash(t *testing.T) {
	msg := PlayerMessage{
		MessageHeader: MessageHeader{
			PrevSignature: nil,
			Sender:        uuid.MustParse("58f6356e-b30c-4811-8bfc-d72a9ee99e73"),
		},
		MessageSignature: fromBase64("L+gR1hrubxVdhPAe6J+nQCH4U+jsUvJDJE+dzsL8DJwpHfUwT4dgSwm7mI/u8rVxrjPVial1DU0sPZodaGVqcSqApyK38bJThWpmojYtieT63hcsgAZzhGNG0GUrykEzHIMPvAnO0bEBmOqPKWjp/kDxhUgF1kKgmQmiOb1fTazi9dVxGuepVkFXknhp7aZBvQ4oQ94bRY5lTMoWyvNcs3CUpVdeFuWwIVnRAIn+hQQ5rDBvWTgKpFTOuxcCOf6hbtPOO2HZ7TT0rsM1D1LV0R9oHUqlEe4nB0E/vT3GdcplSQTSWc7dDmwTjB+wFeGxrNjFP3FEt3he6a+8Q1svlw=="),
		MessageBody: MessageBody{
			PlainMsg:     "123",
			DecoratedMsg: nil,
			Timestamp:    time.Unix(1669990398, 750000000),
			Salt:         -5503869105027681791,
			History:      nil,
		},
		UnsignedContent: nil,
		FilterMask:      FilterMask{},
	}
	hash := msg.MessageBody.Hash()
	want := []byte{
		0x40, 0x2f, 0x63, 0xf1, 0x41, 0x64, 0x83, 0xea,
		0x64, 0xbc, 0xe1, 0xab, 0x4f, 0xa1, 0xdd, 0xcf,
		0x31, 0x6b, 0xdf, 0x9, 0xd3, 0xe3, 0x0, 0xed,
		0xd9, 0x9d, 0x61, 0xb2, 0xfe, 0xe1, 0x23, 0x38,
	}
	if !bytes.Equal(hash, want) {
		t.Errorf("hash not match: want %v, got: %v", want, hash)
	}
	N := new(big.Int)
	N.SetString("19764335557512872060688171398766153113124870942516110890430525316014258628424563821100465499350547856280308462415364939681918846480558609475927282823778386281239015421014618416397562105532153533222767841904849714784545929150813290491806771374240538006775842793512508846836865954304676946257326405255232806869270545112507181469911882584371001804731679283550070980294409246775152428394531383836676761395396452034288295355452943368136472112108083541314359681019016586668426966809565492641015821429276348346645505643438759550398087660394335140584223095644803819139358171070866744685891950091018380254474327892983538632197", 10)
	if err := rsa.VerifyPKCS1v15(&rsa.PublicKey{N: N, E: 65537}, crypto.SHA256, msg.Hash(), msg.MessageSignature); err != nil {
		t.Error(err)
	}
}

func fromBase64(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}
