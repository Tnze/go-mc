package generateutils

import (
	"strings"
	"unicode"
)

func UpperTheFirst(word string) string {
	runes := []rune(word)
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}

func ToGoTypeName(name string) string {
	name = strings.TrimPrefix(name, "minecraft:")
	words := strings.Split(name, "_")
	for i := range words {
		words[i] = UpperTheFirst(words[i])
	}
	return strings.Join(words, "")
}

func ToFuncReceiverName(name string) string {
	if len(name) > 0 {
		name = string(unicode.ToLower([]rune(name)[0]))
	}
	return name
}
