// Usage: go run examples/ping/ping.go localhost
package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/google/uuid"
)

type status struct {
	Description chat.Message
	Players     struct {
		Max    int
		Online int
		Sample []struct {
			ID   uuid.UUID
			Name string
		}
	}
	Version struct {
		Name     string
		Protocol int
	}
	Favicon Icon
	Delay   time.Duration
}

// Icon should be a PNG image that is Base64 encoded
// (without newlines: \n, new lines no longer work since 1.13)
// and prepended with "data:image/png;base64,".
type Icon string

var IconFormatErr = errors.New("data format error")
var IconAbsentErr = errors.New("icon not present")

// ToPNG decode base64-icon, return a PNG image
// Take care of there is no safety check, image may contain malicious code.
func (i Icon) ToPNG() ([]byte, error) {
	const prefix = "data:image/png;base64,"
	if i == "" {
		return nil, IconAbsentErr
	}
	if !strings.HasPrefix(string(i), prefix) {
		return nil, IconFormatErr
	}
	return base64.StdEncoding.DecodeString(strings.TrimPrefix(string(i), prefix))
}

func main() {
	addr := getAddr()
	fmt.Printf("MCPING (%s):", addr)
	resp, delay, err := bot.PingAndList(addr)
	if err != nil {
		fmt.Printf("ping and list server fail: %v", err)
		os.Exit(1)
	}

	var s status
	err = json.Unmarshal(resp, &s)
	if err != nil {
		fmt.Print("unmarshal resp fail:", err)
		os.Exit(1)
	}
	s.Delay = delay

	fmt.Print(&s)
}

func getAddr() string {
	const usage = "Usage: mcping <hostname>[:port]"
	if len(os.Args) < 2 {
		fmt.Println("no host name.", usage)
		os.Exit(1)
	}

	return os.Args[1]
}

var outTemp = template.Must(template.New("output").Parse(`
	Version: [{{ .Version.Protocol }}] {{ .Version.Name }}
	Description: 
{{ .Description }}
	Delay: {{ .Delay }}
	Players: {{ .Players.Online }}/{{ .Players.Max }}{{ range .Players.Sample }}
	- [{{ .Name }}] {{ .ID }}{{ end }}
`))

func (s *status) String() string {
	var sb strings.Builder
	err := outTemp.Execute(&sb, s)
	if err != nil {
		panic(err)
	}
	return sb.String()
}
