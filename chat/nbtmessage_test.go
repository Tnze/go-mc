package chat

import (
	"github.com/Tnze/go-mc/nbt"
	"os"
	"testing"
)

func TestNbtExtraText(t *testing.T) {
	//SNBT: {extra: [{extra: [{color: "dark_gray",text: "> "},{: "test"}],text: ""}],text: ""}
	f, _ := os.Open("testdata/chat.nbt")
	d := nbt.NewDecoder(f)
	var m Message
	if _, err := d.Decode(&m); err != nil {
		t.Fatal(err)
	}

	if m.ClearString() != "> test" {
		t.Fatalf("gets %q, wants %q", m.ClearString(), "> test")
	}
}
