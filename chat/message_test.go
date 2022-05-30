package chat_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/Tnze/go-mc/chat"
	en_us "github.com/Tnze/go-mc/data/lang/en-us"
	pk "github.com/Tnze/go-mc/net/packet"
)

var jsons = []string{
	`{"extra":[{"color":"green","text":"故我依然"},{"color":"white","text":"™ "},{"color":"gray","text":"Kun_QwQ"},{"color":"white","text":": 为什么想要用炼药锅灭火时总是跳不进去"}],"text":""}`,

	`{"translate":"chat.type.text","with":[{"insertion":"Xi_Xi_Mi","clickEvent":{"action":"suggest_command","value":"/tell Xi_Xi_Mi "},"hoverEvent":{"action":"show_entity","value":{"text":"{name:\"{\\\"text\\\":\\\"Xi_Xi_Mi\\\"}\",id:\"c1445a67-7551-4d7e-813d-65ef170ae51f\",type:\"minecraft:player\"}"}},"text":"Xi_Xi_Mi"},"好像是这个id。。"]}`,
	`{"translate":"translation.test.none"}`,
	`{"translate":"translation.test.escape","with":["str1","str2"]}`,
	`{"translate":"translation.test.args","with":["str1","str2"]}`,
	`{"translate":"translation.test.world"}`,

	`"Tnze"`,
	`"§0Tnze"`,
	`"§list"`,

	`{"extra":[" "],"text":""}`,
}

var texts = []string{
	"\033[92m故我依然\033[0m\033[97m™ \033[0m\033[37mKun_QwQ\033[0m\033[97m: 为什么想要用炼药锅灭火时总是跳不进去\033[0m",

	"<Xi_Xi_Mi> 好像是这个id。。",
	"Hello, world!",
	"%s %str1 %%s %%str2",
	"str1 str2",
	"world",

	"Tnze",
	"\033[30mTnze\033[0m",
	"\033[1mist\033[0m",

	" ",
}

var clearTexts = []string{
	"故我依然™ Kun_QwQ: 为什么想要用炼药锅灭火时总是跳不进去",

	"<Xi_Xi_Mi> 好像是这个id。。",
	"Hello, world!",
	"%s %str1 %%s %%str2",
	"str1 str2",
	"world",

	"Tnze",
	"Tnze",
	"ist",

	" ",
}

func TestMessage_String(t *testing.T) {
	chat.SetLanguage(en_us.Map)
	for i, v := range jsons {
		var cm chat.Message
		err := cm.UnmarshalJSON([]byte(v))
		if err != nil {
			t.Error(err)
		}
		if str := cm.String(); str != texts[i] {
			t.Errorf("gets %q, wants %q", str, texts[i])
		}
	}
}

func TestMessage_ClearString(t *testing.T) {
	chat.SetLanguage(en_us.Map)
	for i, v := range jsons {
		var cm chat.Message
		err := json.Unmarshal([]byte(v), &cm)
		if err != nil {
			t.Error(err)
		}

		str := cm.ClearString()
		if str != clearTexts[i] {
			t.Errorf("gets %q, wants %q", str, texts[i])
		}
	}
}

func TestMessage_WriteTo(t *testing.T) {
	chat.SetLanguage(en_us.Map)
	var codeMsg bytes.Buffer
	_, _ = chat.Message{Translate: "multiplayer.disconnect.server_full"}.WriteTo(&codeMsg)

	var msg pk.String // Decode as a String
	if _, err := msg.ReadFrom(&codeMsg); err != nil {
		t.Errorf("decode message fail: %v", err)
	}

	wantMsg := `{"translate":"multiplayer.disconnect.server_full"}`
	if string(msg) != wantMsg {
		t.Error("encode Message error: get", string(msg), ", want", wantMsg)
	}
}

func ExampleMessage_Append() {
	msg := chat.Message{Text: "1111"}.
		Append(chat.Message{Text: "22222"}).
		Append(chat.Message{Text: "333333"}).
		Append(chat.Message{Text: "4444444"})
	fmt.Print(msg)
	// Output: 1111222223333334444444
}

func TestMessage_Append_issue148(t *testing.T) {
	msg := chat.Text("hello").Append(chat.Text("world"))
	if len(msg.Extra) != 1 {
		t.Fatalf("Length of msg.Extra should be 1: %#v", msg)
	}
}

func TestMessage_MarshalJSON_issue151(t *testing.T) {
	// The "text" field should be omitted when "translate" exist.
	// And NOT omitted when "extra" so.
	// That is, we should correctly generate these chat messages:
	// {"text":"","extra":["str1", "str2"]} and
	// {"translate":"translation.test.escape","with":["str1","str2"]}
	mustJson := func(msg *chat.Message) string {
		data, err := json.Marshal(msg)
		if err != nil {
			t.Fatal(err)
		}
		return string(data)
	}
	// "" + "Hello, world!"
	msg1 := chat.Text("").Append(chat.TranslateMsg("translation.test.none"))
	if msg := mustJson(&msg1); !strings.Contains(msg, "text") {
		t.Fatalf("%s doesn't contain %s", msg, "text")
	}
	//
	msg2 := chat.TranslateMsg("translation.test.none")
	if msg := mustJson(&msg2); strings.Contains(msg, "text") {
		t.Fatalf("%s contains %s", msg, "text")
	}
}

func TestMessage_UnmarshalJSON_hoverEvent(t *testing.T) {
	msg := `{
		"text": "Text0",
		"extra": [
			{
				"hoverEvent": {"action": "show_text","value": "la"},
				"text": "Text1"
			},
			{
				"hoverEvent": {"action": "show_text","value": {"text":"la","color":"red"}},
				"text": "Text2"
			},
			{
				"hoverEvent": {"action": "show_text","value": [{"color": "white","text": "Normal"}]},
				"text": "Text3"
			}
		]
	}`
	var message chat.Message
	err := json.Unmarshal([]byte(msg), &message)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(message)
}

func ExampleTranslateMsg() {
	fmt.Println(chat.TranslateMsg("translation.test.none"))
	fmt.Println(chat.TranslateMsg(
		// translation.test.complex == "Prefix, %s%[2]s again %s and %[1]s lastly %s and also %[1]s again!"
		"translation.test.complex",
		chat.Text("1111"),
		chat.Text("2222"),
		chat.Text("3333"),
	).String())
	// Output:
	// Hello, world!
	// Prefix, 11112222 again 3333 and 1111 lastly 2222 and also 1111 again!
}
