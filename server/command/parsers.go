package command

import (
	"io"
	"strconv"
	"strings"

	pk "github.com/Tnze/go-mc/net/packet"
)

type StringParser struct {
	Format int32
	V      string
}

func (s StringParser) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Identifier("brigadier:string"),
		pk.VarInt(s.Format),
	}.WriteTo(w)
}

func (s StringParser) Parse(cmd string) (token string, value interface{}, err error) {
	switch s.Format {
	case 2: // Greedy Phrase
		return cmd, cmd, nil
	case 1: // Quotable Phrase
		if len(cmd) > 0 && cmd[0] == '"' {
			var sb strings.Builder
			var isEscaping bool
			for i, v := range cmd[1:] {
				if isEscaping {
					isEscaping = false
					switch v {
					case '\\':
						sb.WriteRune('\\')
					case '"':
						sb.WriteRune('"')
					}
				} else if v == '\\' {
					isEscaping = true
				} else if v == '"' {
					return cmd[:i], sb.String(), nil
				} else {
					sb.WriteRune(v)
				}
			}
			return "", nil, ParseErr{
				Pos: len(cmd) - 1,
				Err: "expected '\"'",
			}
		}
		fallthrough
	case 0: // Single Word
		i := strings.IndexAny(cmd, "\t\n\v\f\r ")
		if i == -1 {
			return cmd, cmd, nil
		}
		return cmd[:i], cmd[:i], nil
	default:
		panic("StringParser: unknown format 0x" + strconv.FormatInt(int64(s.Format), 16))
	}
}

type ParseErr struct {
	Pos int
	Err string
}

func (p ParseErr) Error() string {
	return p.Err
}
