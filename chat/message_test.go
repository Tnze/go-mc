package chat_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Tnze/go-mc/chat"
	en_us "github.com/Tnze/go-mc/data/lang/en-us"
	pk "github.com/Tnze/go-mc/net/packet"
)

/*
	"translation.test.none":                                      "Hello, world!",
	"translation.test.complex":                                   "Prefix, %s %[2]s again %s   and %[1]s lastly %s   and also %[1]s again!"
																  "Prefix, str1str2 again str3 and str1  lastly str2 and also str1  again!"
																  "Prefix, str1str2 again str2 and str1  lastly str3 and also str1  again!"
	"translation.test.escape":                                    "%%s %%%s %%%%s %%%%%s",
	"translation.test.invalid":                                   "hi %",
	"translation.test.invalid2":                                  "hi %  s",
	"translation.test.args":                                      "%s %s",
	"translation.test.world":
*/
var jsons = []string{
	`{"extra":[{"color":"green","text":"故我依然"},{"color":"white","text":"™ "},{"color":"gray","text":"Kun_QwQ"},{"color":"white","text":": 为什么想要用炼药锅灭火时总是跳不进去"}],"text":""}`,

	`{"translate":"chat.type.text","with":[{"insertion":"Xi_Xi_Mi","clickEvent":{"action":"suggest_command","value":"/tell Xi_Xi_Mi "},"hoverEvent":{"action":"show_entity","value":{"text":"{name:\"{\\\"text\\\":\\\"Xi_Xi_Mi\\\"}\",id:\"c1445a67-7551-4d7e-813d-65ef170ae51f\",type:\"minecraft:player\"}"}},"text":"Xi_Xi_Mi"},"好像是这个id。。"]}`,
	`{"translate":"translation.test.none"}`,
	//`{"translate":"translation.test.complex","with":["str1","str2","str3"]}`,
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
	//"Prefix, str1str2 again str2 and str1 lastly str3 and also str1 again!",
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
	//"Prefix, str1str2 again str2 and str1 lastly str3 and also str1 again!",
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
		err := cm.UnmarshalJSON([]byte(v))
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
	_, _ = chat.TranslateMsg("multiplayer.disconnect.server_full").WriteTo(&codeMsg)

	var msg pk.String // Decode as a String
	if _, err := msg.ReadFrom(&codeMsg); err != nil {
		t.Errorf("decode message fail: %v", err)
	}

	wantMsg := `{"translate":"multiplayer.disconnect.server_full"}`
	if string(msg) != wantMsg {
		t.Error("encode Message error: get", string(msg), ", want", wantMsg)
	}
}

func TestMessage_MarshalJSON(t *testing.T) {
	messages := []chat.Message{
		chat.Text("Hello World!"),
		chat.TranslateMsg("example.key"),
		chat.TranslateMsg("example.key", chat.Text("another message"), chat.TranslateMsg("another.key")),
		{MessageText: &chat.MessageText{Text: "123"}, Extra: []chat.Message{chat.TranslateMsg("example.key", chat.Text("another message"))}},
		{Extra: []chat.Message{chat.Text("123")}},
	}

	jsons := []string{
		`{"text":"Hello World!"}`,
		`{"translate":"example.key"}`,
		`{"translate":"example.key","with":[{"text":"another message"},{"translate":"another.key"}]}`,
		`{"text":"123","extra":[{"translate":"example.key","with":[{"text":"another message"}]}]}`,
		`{"text":"","extra":[{"text":"123"}]}`,
	}

	for i, m := range messages {
		result, err := json.Marshal(m)
		if err != nil {
			t.Errorf("marshal Message fail: %v", err)
			return
		}

		wantMsg := jsons[i]
		if string(result) != wantMsg {
			t.Error("Message JSON marshal error: get", string(result), ", want", wantMsg)
		}
	}
}

func TestMessage_UnmarshalJSON(t *testing.T) {
	jsons := []string{
		`{"text":"Hello World!"}`,
		`{"translate":"example.key"}`,
		`{"translate":"example.key","with":[{"text":"another message"},{"translate":"another.key"}]}`,
		`{"text":"123","extra":[{"translate":"example.key","with":[{"text":"another message"}]}]}`,
		`{"text":"","extra":[{"text":"123"}]}`,
	}

	type test func(message chat.Message) bool
	tests := []test{
		func(message chat.Message) bool {
			return message.MessageTranslation == nil && message.Text == "Hello World!" && message.GetText() == "Hello World!"
		},
		func(message chat.Message) bool {
			return message.MessageText == nil && message.GetText() == "" && message.Translate == "example.key"
		},
		func(message chat.Message) bool {
			return message.MessageText == nil && message.Translate == "example.key" && len(message.With) == 2 &&
				bytes.Equal(message.With[0], []byte(`{"text":"another message"}`)) &&
				bytes.Equal(message.With[1], []byte(`{"translate":"another.key"}`))
		},
		func(message chat.Message) bool {
			extra := message.Extra[0]
			return message.MessageTranslation == nil && message.GetText() == "123" && len(message.Extra) == 1 &&
				extra.MessageText == nil && len(extra.With) == 1 &&
				bytes.Equal(extra.With[0], []byte(`{"text":"another message"}`))
		},
		func(message chat.Message) bool {
			return message.MessageTranslation == nil && message.Text == "" && message.GetText() == "" &&
				len(message.Extra) == 1 && message.Extra[0].GetText() == "123"
		},
	}

	for i, j := range jsons {
		var msg chat.Message
		err := json.Unmarshal([]byte(j), &msg)
		if err != nil {
			t.Errorf("unmarshal Message fail: %v", err)
			return
		}

		if !tests[i](msg) {
			t.Errorf("test for json `%s` failed, got %#v", j, msg)
		}
	}
}

func ExampleMessage_Append() {
	msg := chat.Text("1111").
		Append(chat.Text("22222").
			Append(chat.Text("333333")).
			Append(chat.Text("4444444")))
	fmt.Print(msg)
	// Output: 1111222223333334444444
}

func TestMessage_Append_issue148(t *testing.T) {
	msg := chat.Text("hello").Append(chat.Text("world"))
	if len(msg.Extra) != 1 {
		t.Fatalf("Length of msg.Extra should be 1: %#v", msg)
	}
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
